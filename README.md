<p align="center">
  <img src="https://img.shields.io/badge/version-1.0.0-f43f5e?style=for-the-badge" alt="version">
  <img src="https://img.shields.io/badge/go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="go">
  <img src="https://img.shields.io/badge/license-MIT-green?style=for-the-badge" alt="license">
  <img src="https://img.shields.io/badge/platform-Windows%20%7C%20Linux%20%7C%20macOS-blue?style=for-the-badge" alt="platform">
</p>

<h1 align="center">YT Audio Downloader</h1>

<p align="center">
  Download YouTube audio as high-quality MP3 files.<br>
  Fast, clean, and easy to use.
</p>

---

## Features

- **One-click download** — paste a YouTube link, get an MP3
- **Best quality** — automatically picks the highest audio quality available
- **Folder picker** — native OS dialog to choose where to save
- **Real-time progress** — watch the download and conversion live
- **Cross-platform** — works on Windows, Linux, and macOS
- **Zero config** — auto-finds `yt-dlp` and `ffmpeg` in PATH or common install locations
- **Dark UI** — clean, modern interface

## Demo

```
1. Paste a YouTube URL
2. Click "Change Folder" to pick where to save
3. Click "Download Audio"
4. Done — MP3 is in your chosen folder
```

## Requirements

| Tool | Purpose | Install |
|------|---------|---------|
| [yt-dlp](https://github.com/yt-dlp/yt-dlp) | Download engine | `pip install yt-dlp` |
| [ffmpeg](https://ffmpeg.org/) | Audio conversion | `winget install ffmpeg` |

## Install

### Download (Windows)

1. Go to [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases/latest)
2. Download `YT Audio Downloader.exe`
3. Place it anywhere and double-click to run

### Build from Source

```bash
git clone https://github.com/ujanchakraborty/yt-audio-downloader.git
cd yt-audio-downloader/app

# Build
go build -ldflags="-s -w" -o yt-audio-downloader .

# Run
./yt-audio-downloader
```

## Usage

1. Double-click the `.exe` (or run the binary)
2. Your browser opens automatically with the app
3. Paste any YouTube URL
4. Click **Change Folder** to pick a save location
5. Click **Download Audio**
6. Watch the progress — your MP3 appears in the chosen folder

## Project Structure

```
yt-audio-downloader/
├── app/
│   ├── main.go          # Server + download logic
│   ├── go.mod
│   └── ui/
│       └── index.html   # Frontend (embedded in binary)
├── .github/
│   └── workflows/
│       └── release.yml  # Auto-build on tag
├── .gitignore
├── LICENSE
└── README.md
```

## How It Works

- A small Go web server starts on `localhost` and opens your browser
- The frontend sends YouTube URLs to the backend
- The backend calls `yt-dlp` to download the audio and convert to MP3
- Progress is streamed back to the UI in real-time
- The binary embeds the frontend — single file, no dependencies

## License

[MIT](LICENSE)

---

<p align="center">
  Built with Go + yt-dlp
</p>
