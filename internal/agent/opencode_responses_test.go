package agent

import "testing"

// a gpt-6+ model is OpenCode's on Responses, where an effort change keeps
// the cache; any other stays on the provider's Chat (#3145)
func TestOpenCodeResponses(t *testing.T) {
	for id, want := range map[string]bool{
		"codex/gpt-6.1-sol":  true,
		"codex/gpt-6-luna":   true,
		"gpt-7":              true,
		"relay/GPT-10.2-max": true,
		"codex/gpt-5.5":      false,
		"gpt-4o":             false,
		"relay/glm-4.6":      false,
		"claude-sonnet-4-5":  false,
		"my-gpt-6x":          false,
	} {
		if got := openCodeResponses(id); got != want {
			t.Errorf("%s: %v, want %v", id, got, want)
		}
	}
}
