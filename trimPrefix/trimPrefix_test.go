package trimPrefix

import (
	"testing"
)

func TestTrimPrefix(t *testing.T) {
	s := "Greetings Soldier"
	trimmed := trimPrefix(s, "Greetings ")

	expected := "Soldier"
	if trimmed != expected {
		t.Errorf("expected %q, got %q", expected, trimmed)
	}
}
