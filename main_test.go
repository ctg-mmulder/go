package main

import (
	"testing"
)

func TestRoundToNearest50(t *testing.T) {
	t.Run("position 101, 101 should return 100, 100", func(t *testing.T) {
		if got := RoundUpToNearest50(101); got != 100 {
			t.Errorf("is not hunderd")
		}
	})
	t.Run("position 141 should return 150", func(t *testing.T) {
		if got := RoundUpToNearest50(141); got != 150 {
			t.Errorf("is not hunderd")
		}
	})
	t.Run("position 0 should return 50", func(t *testing.T) {
		if got := RoundUpToNearest50(0); got != 50 {
			t.Errorf("is not fifty")
		}
	})
	t.Run("position 0 should return 50", func(t *testing.T) {
		if got := RoundUpToNearest50(0); got != 50 {
			t.Errorf("is not fifty")
		}
	})
}
