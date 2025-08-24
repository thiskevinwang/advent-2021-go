package main

import (
	"testing"
)

// test the algorithm
func TestDay7(t *testing.T) {
	t.Parallel()

	got := Day7([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14})

	if got != 2 {
		t.Errorf("Day7(16, 1, 2, 0, 4, 2, 7, 1, 2, 14) = %d; want 2", got)
	}
}

func TestDay7_2(t *testing.T) {
	t.Parallel()

	got := Day7([]int{1, 1, 1, 1, 1, 1})

	if got != 1 {
		t.Errorf("Day7() = %d; want 1", got)
	}
}

func TestDay7_3(t *testing.T) {
	t.Parallel()

	got := Day7([]int{0, 1})

	if got != 1 {
		t.Errorf("Day7() = %d; want 1", got)
	}
}
