package game

import "testing"

func TestPlay(t *testing.T) {
	x := play()
	if x != 0 {
		t.Errorf("Expected 0; but got %d", x)
	}
}
