# Discord Cleanup CLI

A minimal Go-based command-line utility that periodically cleans up stale messages posted by the bot itself in specified Discord channels.

## Features
- Reads configuration from environment variables.
- Uses the official `discordgo` library.
- Runs on a minimal distroless container.
- Configurable cleanup age and polling interval.
- **Skips pinned messages** to preserve important bot posts.

## Environment Variables
| Variable | Description | Default |
|---|---|---|
| `DISCORD_TOKEN` | Bot token (required) | |
| `DISCORD_CHANNELS` | Comma-separated list of channel IDs to monitor | |
| `CLEANUP_AGE_HOURS` | Age in hours after which a message is considered stale | `24` |
| `POLL_INTERVAL_MINUTES` | How often the cleanup runs in `loop` mode | `360` |
| `MODE` | `oneshot` (default) or `loop` | `oneshot` |

## Local Development

### Build
```bash
go build -o discord-cleanup .
```

### Run (One-Shot)
```bash
export DISCORD_TOKEN="your_token"
export DISCORD_CHANNELS="channel_id_1,channel_id_2"
./discord-cleanup
```

### Run (Loop)
```bash
export MODE=loop
./discord-cleanup
```

### Docker
```bash
make build-dev
docker run --env-file .env discord-cleanup:local-dev
```
## License
MIT © kwv
