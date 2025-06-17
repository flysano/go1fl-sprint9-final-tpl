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

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Int()
	}
	return slice
}

func maximum(data []int) int {
	if len(data) == 0 {
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

func maxChunks(data []int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	sliceSize := len(data) / CHUNKS
	maxNumsSlice := make([]int, CHUNKS)

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		startSlice := i * sliceSize
		endSlice := startSlice + sliceSize
		if i == CHUNKS-1 {
			endSlice = len(data)
		}
		go func(index, start, end int) {
			defer wg.Done()
			mu.Lock() //гонки быть не должно, но для большей надежности добавил мьютекс
			max := maximum(data[start:end])
			mu.Unlock()
			maxNumsSlice[index] = max
		}(i, startSlice, endSlice)
	}
	wg.Wait()

	return maximum(maxNumsSlice)
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
