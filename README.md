# YT Audio Downloader

Download YouTube audio as high-quality MP3.

## Install

### Windows
1. Download `YT Audio Downloader.exe` from [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases/latest)
2. Install [yt-dlp](https://github.com/yt-dlp/yt-dlp) and [ffmpeg](https://ffmpeg.org/) (see below)
3. Double-click the .exe to run

### Linux
1. Download `yt-audio-downloader` and `run.sh` from [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases/latest)
2. Install dependencies (see below)
3. Open terminal in the download folder and run:
   ```bash
   chmod +x run.sh
   ```
   ```bash
   ./run.sh
   ```
   > **Note:** Linux binaries run from the terminal, not by double-clicking.

### macOS
1. Download `yt-audio-downloader-mac` from [Releases](https://github.com/ujanchakraborty/yt-audio-downloader/releases/latest)
2. Install yt-dlp and ffmpeg (see below)
3. Open terminal in the download folder and run:
   ```bash
   chmod +x yt-audio-downloader-mac
   ```
   ```bash
   ./yt-audio-downloader-mac
   ```
   > **Note:** If Gatekeeper blocks it, allow it in System Settings → Privacy & Security.

## Installing Dependencies

The app needs [yt-dlp](https://github.com/yt-dlp/yt-dlp) and [ffmpeg](https://ffmpeg.org/) to work.

**Windows:**
```
pip install yt-dlp
winget install ffmpeg
```

**Linux:**
```bash
sudo wget https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -O /usr/local/bin/yt-dlp && sudo chmod a+rx /usr/local/bin/yt-dlp && sudo apt update && sudo apt install -y ffmpeg
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
