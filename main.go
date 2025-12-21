package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var version = "dev"

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parseChannels(val string) []string {
	parts := strings.Split(strings.TrimSpace(val), ",")
	var out []string
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

func main() {
	log.Printf("Discord Cleanup CLI version: %s\n", version)
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN not set")
	}
	channels := parseChannels(getEnv("DISCORD_CHANNELS", ""))
	if len(channels) == 0 {
		log.Fatal("DISCORD_CHANNELS not set or empty")
	}
	cleanupHours, _ := strconv.Atoi(getEnv("CLEANUP_AGE_HOURS", "24"))
	pollMinutes, _ := strconv.Atoi(getEnv("POLL_INTERVAL_MINUTES", "360"))
	mode := strings.ToLower(getEnv("MODE", "oneshot"))

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
	}
	if err = dg.Open(); err != nil {
		log.Fatalf("error opening connection: %v", err)
	}
	defer func() { _ = dg.Close() }()

	if mode == "oneshot" {
		log.Println("Running in oneshot mode")
		cutoff := time.Now().Add(-time.Duration(cleanupHours) * time.Hour)
		runCleanup(dg, channels, cutoff)
		log.Println("Cleanup complete. Exiting.")
		return
	}

	log.Printf("Running in loop mode (poll interval: %d minutes)\n", pollMinutes)
	cutoff := time.Now().Add(-time.Duration(cleanupHours) * time.Hour)
	ticker := time.NewTicker(time.Duration(pollMinutes) * time.Minute)
	defer ticker.Stop()

	// initial run
	runCleanup(dg, channels, cutoff)

	for range ticker.C {
		cutoff = time.Now().Add(-time.Duration(cleanupHours) * time.Hour)
		runCleanup(dg, channels, cutoff)
	}
}

func runCleanup(dg *discordgo.Session, channels []string, cutoff time.Time) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	for _, chID := range channels {
		msgs, err := dg.ChannelMessages(chID, 100, "", "", "")
		if err != nil {
			log.Printf("failed to fetch messages for channel %s: %v", chID, err)
			continue
		}
		for _, m := range msgs {
			if m.Author.ID != dg.State.User.ID {
				continue // not our bot
			}
			if m.Timestamp.After(cutoff) {
				continue // recent
			}
			if m.Pinned {
				continue // skip pinned messages
			}
			if err := dg.ChannelMessageDelete(chID, m.ID); err != nil {
				log.Printf("failed to delete message %s in %s: %v", m.ID, chID, err)
			} else {
				log.Printf("deleted stale message %s in %s", m.ID, chID)
			}
		}
	}
	select {
	case <-ctx.Done():
	default:
	}
}
