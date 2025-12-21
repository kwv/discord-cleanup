# This Dockerfile is optimized for GoReleaser and fast local development.
# It expects the 'discord-cleanup' binary to be present in the build context.

FROM gcr.io/distroless/static-debian12
COPY discord-cleanup /discord-cleanup
ENTRYPOINT ["/discord-cleanup"]
