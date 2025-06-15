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
	t.Run("Empty slice", func(t *testing.T) {
		slice := make([]int, 0, 0)
		num := maximum(slice)
		assert.Zero(t, num)
	})

	t.Run("One element slice", func(t *testing.T) {
		slice := []int{45}
		num := maximum(slice)
		assert.Equal(t, slice[0], num)
	})

}
