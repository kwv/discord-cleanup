# ---------- Build stage ----------
FROM golang:1.22-alpine AS builder

# VERSION is passed via --build-arg
ARG VERSION=unknown

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w -X main.version=${VERSION}" -o discord-cleanup

# ---------- Runtime stage ----------
FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/discord-cleanup /discord-cleanup
ENTRYPOINT ["/discord-cleanup"]
