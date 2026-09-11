# YT Audio Downloader

Download YouTube audio as high-quality MP3.

## Install

### Windows
1. Download `YT Audio Downloader.exe` from [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases/latest)
2. Install [yt-dlp](https://github.com/yt-dlp/yt-dlp) and [ffmpeg](https://ffmpeg.org/) (see below)
3. Double-click the .exe to run

### Linux
1. Download `yt-audio-downloader` from [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases/latest)
2. Install yt-dlp and ffmpeg (see below)
3. Open terminal in the download folder, then run:
   ```bash
   chmod +x yt-audio-downloader
   ```
   ```bash
   ./yt-audio-downloader
   ```

### macOS
1. Download `yt-audio-downloader-mac` from [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases/latest)
2. Install yt-dlp and ffmpeg (see below)
3. Open terminal in the download folder, then run:
   ```bash
   chmod +x yt-audio-downloader-mac
   ```
   ```bash
   ./yt-audio-downloader-mac
   ```

## Installing Dependencies

The app needs [yt-dlp](https://github.com/yt-dlp/yt-dlp) and [ffmpeg](https://ffmpeg.org/) to work.

**Windows:**
```
pip install yt-dlp
winget install ffmpeg
```

**Linux:**
```
pip install yt-dlp
sudo apt install ffmpeg
```

**macOS:**
```
pip install yt-dlp
brew install ffmpeg
```

## Build from Source

Requires [Go](https://go.dev/dl/) 1.22+.

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
