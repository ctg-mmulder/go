package main

import (
	"testing"
)

func TestIsWhite(t *testing.T) {

	t.Run("Even number should return true", func(t *testing.T) {
		if got := IsWhite(2); got != true {
			t.Errorf("IsWhite(2) = %v; want true", got)
		}
	})
	t.Run("Odd number should return false", func(t *testing.T) {
		if got := IsWhite(3); got != false {
			t.Errorf("IsWhite(3) = %v; want false", got)
		}
	})
}

func TestRoundToNearest50(t *testing.T) {

}
