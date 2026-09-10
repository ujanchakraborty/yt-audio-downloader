# YT Audio Downloader

Download YouTube audio as high-quality MP3 files.

## Features

- Paste any YouTube URL, get the best quality MP3
- Built-in folder picker to choose save location
- Real-time download progress
- Auto-finds `yt-dlp` and `ffmpeg` in PATH or common install locations
- Clean dark UI

## Requirements

- [yt-dlp](https://github.com/yt-dlp/yt-dlp) installed and in PATH
- [ffmpeg](https://ffmpeg.org/) installed and in PATH (for audio conversion)

## Install

Download the latest `.exe` (Windows) from [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases).

## Build from Source

```bash
# Install Go 1.23+
# Then:
cd app
go mod tidy
go build -ldflags="-s -w" -o yt-audio-downloader .
```

## Usage

1. Double-click the `.exe`
2. Browser opens automatically
3. Paste a YouTube URL
4. Click **Change Folder** to pick where to save
5. Click **Download Audio**

## License

MIT
