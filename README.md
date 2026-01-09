# Emoji Smith Bot

A simple Telegram bot that converts images/videos into emoji-ready formats and sends the processed file back to you.

## Requirements

- Go
- FFmpeg
- Telegram bot token

### Install Go

- https://go.dev/dl/

### Install FFmpeg

- https://www.ffmpeg.org/download.html

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
