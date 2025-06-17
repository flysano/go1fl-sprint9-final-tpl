package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("Zero size", func(t *testing.T) {
		slice := generateRandomElements(0)
		assert.Empty(t, slice)
	})

	t.Run("Negative", func(t *testing.T) {
		slice := generateRandomElements(-1)
		assert.Empty(t, slice)
	})

	t.Run("OK size", func(t *testing.T) {
		size := 34
		slice := generateRandomElements(size)
		assert.Equal(t, size, len(slice))
	})
}

func TestMaximun(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{32}, 32},
		{"negative number", []int{-3, -4, -1}, -1},
		{"mixed", []int{-5, 6, 19, 0, -45}, 19},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			v := maximum(tst.input)
			assert.Equal(t, tst.expected, v)
		})
	}
}
