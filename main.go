package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		fmt.Printf("Slice creation error with size %d\n", size)
		return []int{}
	}
	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Int()
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 || len(data) == 1 {
		fmt.Printf("size")
		return 0
	}
	max := data[0]
	for _, num := range data[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	sliceSize := len(data) / CHUNKS
	maxSlice := make([]int, CHUNKS)
	var wg sync.WaitGroup

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		go func(i int) {
			defer wg.Done()

			startSlice := i * sliceSize
			endSlice := startSlice + sliceSize

			if i == CHUNKS-1 {
				endSlice = len(data)
			}
			max := data[startSlice]
			for _, num := range data[startSlice:endSlice] {
				if num > max {
					max = num
				}
			}
			maxSlice[i] = max
		}(i)
	}
	wg.Wait()

	endMax := maxSlice[0]
	for _, num := range maxSlice[1:] {
		if num > endMax {
			endMax = num
		}
	}
	return endMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	sliceNums := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(sliceNums)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Milliseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(sliceNums)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Milliseconds())
}
