package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:embed ui/*
var staticFiles embed.FS

var version = "dev"

type downloadReq struct {
	URL string `json:"url"`
	Dir string `json:"dir"`
}

type resolveReq struct {
	Name string `json:"name"`
}

type dirReq struct {
	Dir string `json:"dir"`
}

type logEntry struct {
	Text string `json:"text"`
	Type string `json:"type"`
	Time string `json:"time"`
}

var (
	mu        sync.Mutex
	logs      []logEntry
	busy      bool
	outputDir string
)

func main() {
	home, _ := os.UserHomeDir()
	outputDir = filepath.Join(home, "Downloads")

	mux := http.NewServeMux()

	uiSub, _ := fs.Sub(staticFiles, "ui")
	mux.Handle("/", http.FileServer(http.FS(uiSub)))

	mux.HandleFunc("GET /api/info", apiInfo)
	mux.HandleFunc("POST /api/setdir", apiSetDir)
	mux.HandleFunc("POST /api/resolve", apiResolve)
	mux.HandleFunc("POST /api/download", apiDownload)
	mux.HandleFunc("GET /api/logs", apiLogs)
	mux.HandleFunc("POST /api/logs/clear", apiClearLogs)

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		fmt.Println("Press Enter to exit...")
		fmt.Scanln()
		os.Exit(1)
	}

	port := listener.Addr().(*net.TCPAddr).Port
	url := fmt.Sprintf("http://127.0.0.1:%d", port)

	fmt.Println()
	fmt.Println("  ╭─────────────────────────╮")
	fmt.Println("  │   YT Audio Downloader   │")
	fmt.Printf("  │   v%-20s│\n", version)
	fmt.Println("  ╰─────────────────────────╯")
	fmt.Println()
	fmt.Printf("  Open: %s\n", url)
	fmt.Println("  Close this window to stop.")
	fmt.Println()

	openBrowser(url)
	http.Serve(listener, mux)
}

func apiInfo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(map[string]string{
		"dir":     outputDir,
		"version": version,
	})
}

func apiSetDir(w http.ResponseWriter, r *http.Request) {
	var req dirReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Dir == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if info, err := os.Stat(req.Dir); err != nil || !info.IsDir() {
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid directory"})
		return
	}
	mu.Lock()
	outputDir = req.Dir
	mu.Unlock()
	json.NewEncoder(w).Encode(map[string]string{"dir": req.Dir})
}

func apiResolve(w http.ResponseWriter, r *http.Request) {
	var req resolveReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	home, _ := os.UserHomeDir()
	candidates := []string{
		filepath.Join(home, "Downloads", req.Name),
		filepath.Join(home, "Desktop", req.Name),
		filepath.Join(home, "Documents", req.Name),
		filepath.Join(home, "Music", req.Name),
		filepath.Join(home, "Videos", req.Name),
		filepath.Join(home, "Pictures", req.Name),
		filepath.Join(home, req.Name),
	}

	if runtime.GOOS == "windows" {
		for _, drive := range []string{"C", "D", "E"} {
			candidates = append(candidates, filepath.Join(drive+":\\", req.Name))
		}
	} else {
		candidates = append(candidates, filepath.Join("/", req.Name))
	}

	for _, path := range candidates {
		if fi, err := os.Stat(path); err == nil && fi.IsDir() {
			json.NewEncoder(w).Encode(map[string]string{"path": path})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
}

func apiDownload(w http.ResponseWriter, r *http.Request) {
	var req downloadReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	if busy {
		mu.Unlock()
		json.NewEncoder(w).Encode(map[string]string{"error": "busy"})
		return
	}
	busy = true
	dir := outputDir
	if req.Dir != "" {
		if fi, err := os.Stat(req.Dir); err == nil && fi.IsDir() {
			dir = req.Dir
			outputDir = dir
		}
	}
	mu.Unlock()

	json.NewEncoder(w).Encode(map[string]string{"status": "started"})
	go doDownload(req.URL, dir)
}

func apiLogs(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func apiClearLogs(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	logs = nil
	mu.Unlock()
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func doDownload(url, dir string) {
	defer func() {
		mu.Lock()
		busy = false
		mu.Unlock()
	}()

	addLog("Starting download...", "accent")

	ytdl := findBin("yt-dlp")
	if ytdl == "" {
		addLog("yt-dlp not found. Install: pip install yt-dlp", "error")
		return
	}
	addLog("yt-dlp: "+filepath.Base(ytdl), "dim")

	ffmpegDir := findBin("ffmpeg")
	if ffmpegDir != "" {
		ffmpegDir = filepath.Dir(ffmpegDir)
		addLog("ffmpeg: found", "dim")
	} else {
		addLog("ffmpeg not found — audio conversion may fail", "warning")
	}

	addLog("Save: "+dir, "dim")

	out := filepath.Join(dir, "%(title)s.%(ext)s")
	args := []string{
		"-x", "--audio-format", "mp3",
		"--audio-quality", "0",
		"-o", out,
		"--no-warnings", "--no-check-certificates",
		"--newline", "--progress",
	}
	if ffmpegDir != "" {
		args = append(args, "--ffmpeg-location", ffmpegDir)
	}
	args = append(args, url)

	addLog("Fetching video info...", "info")

	cmd := exec.Command(ytdl, args...)
	cmd.Dir = dir
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	output := stdout.String() + stderr.String()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() != 0 {
			msg := strings.TrimSpace(stderr.String())
			if msg == "" {
				msg = err.Error()
			}
			if len(msg) > 500 {
				msg = msg[:500]
			}
			addLog("Error: "+msg, "error")
			return
		}
	}

	var downloadedFile string
	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "[download]") {
			if strings.Contains(line, "Destination:") {
				idx := strings.Index(line, "Destination:")
				downloadedFile = strings.TrimSpace(line[idx+12:])
				addLog("Saving: "+filepath.Base(downloadedFile), "dim")
			} else if strings.Contains(line, "100%") {
				addLog("Download complete!", "success")
			} else if pct := extractPct(line); pct > 0 {
				addLog(fmt.Sprintf("Progress: %.0f%%", pct), "dim")
			}
		} else if strings.Contains(line, "[ExtractAudio]") {
			addLog("Converting to MP3...", "warning")
		}
	}

	if downloadedFile != "" {
		base := strings.TrimSuffix(strings.TrimSuffix(filepath.Base(downloadedFile), ".webm"), ".m4a")
		mp3 := filepath.Join(dir, base+".mp3")
		if fi, err := os.Stat(mp3); err == nil {
			mb := float64(fi.Size()) / 1024 / 1024
			addLog(fmt.Sprintf("Saved: %s (%.1f MB)", filepath.Base(mp3), mb), "success")
		}
	}

	addLog("Done!", "success")
}

func findBin(name string) string {
	if p, err := exec.LookPath(name); err == nil {
		return p
	}
	if runtime.GOOS == "windows" {
		exe, _ := os.Executable()
		exeDir := filepath.Dir(exe)
		if _, err := os.Stat(filepath.Join(exeDir, name+".exe")); err == nil {
			return filepath.Join(exeDir, name+".exe")
		}
		home, _ := os.UserHomeDir()
		for _, sub := range []string{
			filepath.Join("AppData", "Local", "Microsoft", "WinGet", "Packages"),
			filepath.Join("AppData", "Local", "Programs", "Python"),
			filepath.Join("AppData", "Roaming", "Python"),
		} {
			matches, _ := filepath.Glob(filepath.Join(home, sub, "*"))
			for _, m := range matches {
				p := filepath.Join(m, name+".exe")
				if _, err := os.Stat(p); err == nil {
					return p
				}
			}
		}
	}
	return ""
}

func extractPct(line string) float64 {
	re := regexp.MustCompile(`(\d+\.?\d*)%`)
	m := re.FindString(line)
	if m == "" {
		return 0
	}
	v, _ := strconv.ParseFloat(strings.TrimSuffix(m, "%"), 64)
	return v
}

func addLog(text, typ string) {
	mu.Lock()
	defer mu.Unlock()
	logs = append(logs, logEntry{
		Text: text, Type: typ,
		Time: time.Now().Format("15:04:05"),
	})
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", strings.ReplaceAll(url, "&", "^&"))
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	go cmd.Run()
}
