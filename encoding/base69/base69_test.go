package base69

import (
	"testing"
)

// TestPlaceholder to make sure things are running
func TestPlaceholder(t *testing.T) {
	hello := Placeholder()
	if hello != "ohai!" {
		t.Error("Placeholder should say ohai!")
	}
}
