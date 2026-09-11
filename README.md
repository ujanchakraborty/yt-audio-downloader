# YT Audio Downloader

Download YouTube audio as high-quality MP3.

## Requirements

- [yt-dlp](https://github.com/yt-dlp/yt-dlp) — `pip install yt-dlp`
- [ffmpeg](https://ffmpeg.org/) — `winget install ffmpeg` (Windows) or `sudo apt install ffmpeg` (Linux)

## Install

### Windows
1. Download `YT Audio Downloader.exe` from [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases/latest)
2. Double-click to run

### Linux
1. Download `yt-audio-downloader` from [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases/latest)
2. Make it executable: `chmod +x yt-audio-downloader`
3. Run: `./yt-audio-downloader`

### macOS
1. Download `yt-audio-downloader-mac` from [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases/latest)
2. Make it executable: `chmod +x yt-audio-downloader-mac`
3. Run: `./yt-audio-downloader-mac`

### Build from Source
```bash
git clone https://github.com/ujanchakraborty/yt-audio-downloader.git
cd yt-audio-downloader/app
go build -o yt-audio-downloader .
```

## Usage

1. Run the app — browser opens automatically
2. Paste a YouTube URL
3. Click **Change Folder** to pick where to save
4. Click **Download Audio**

## License

MIT
