package main

import (
	"testing"
)

func TestParseChannels(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"123,456", []string{"123", "456"}},
		{" 123 ,  456  ", []string{"123", "456"}},
		{"123", []string{"123"}},
		{"", nil},
		{" , ", nil},
	}

	for _, tc := range tests {
		got := parseChannels(tc.input)
		if len(got) != len(tc.expected) {
			t.Errorf("parseChannels(%q) = %v; want %v", tc.input, got, tc.expected)
			continue
		}
		for i := range got {
			if got[i] != tc.expected[i] {
				t.Errorf("parseChannels(%q)[%d] = %q; want %q", tc.input, i, got[i], tc.expected[i])
			}
		}
	}
}

func TestGetEnv(t *testing.T) {
	key := "TEST_DISCORD_CLEANUP_ENV"
	fallback := "default"
	
	if got := getEnv(key, fallback); got != fallback {
		t.Errorf("getEnv(%q, %q) = %q; want %q", key, fallback, got, fallback)
	}
}
