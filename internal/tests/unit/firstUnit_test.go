package unit

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestAddUnit(t *testing.T) {
	result := add(2, 3)
	want := 5

	if result != want {
		t.Errorf(`add(2,3) = %d, want match for %d`, result, want)
	}
}

// go test ./tests/... // Runs all test in the folder test.
