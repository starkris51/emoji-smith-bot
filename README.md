# Emoji Smith Bot

A simple Telegram bot that converts images/videos into emoji-ready formats and sends the processed file back to you.

## Requirements

- Go
- FFmpeg
- Telegram bot token

### Install Go

- https://go.dev/dl/

### Install FFmpeg

- Ubuntu/Debian:
  ```sh
  sudo apt update && sudo apt install -y ffmpeg
  ```
- Arch Linux:
  ```sh
  sudo pacman -S ffmpeg
  ```
- macOS (Homebrew):
  ```sh
  brew install ffmpeg
  ```
- Windows:
  - Install FFmpeg from https://ffmpeg.org/download.html and ensure `ffmpeg` is on your `PATH`.

## Setup

Create a `.env` file in the repo root:

```env
TOKEN=123456789:your-telegram-bot-token-here
```

## Run the bot (local)

From the repository root:

```sh
go run ./cmd/bot
```
