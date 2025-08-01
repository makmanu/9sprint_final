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

var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	list := make([]int, 0, size)
	if size < 1 {
		return []int {}
	}
	for range size {
		list = append(list, rand.Int())
	}
	return list
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) < 1 {
		return 0
	}
	max := 0
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) < CHUNKS {
		return maximum(data)
	}
	resultArray := make([]int, CHUNKS)
	sizeOfChunk := len(data) / CHUNKS
	for i := range CHUNKS {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var chunk []int
			if i != CHUNKS - 1 {
				chunk = data[i * sizeOfChunk:(i + 1) * sizeOfChunk]
			} else {
				chunk = data[i * sizeOfChunk:]
			}
			max := maximum(chunk)
			resultArray[i] = max
		}(i)
	}
	wg.Wait()
	return maximum(resultArray)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	randomIntSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	t := time.Now()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maximum(randomIntSlice), time.Since(t).Milliseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	
	t = time.Now()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxChunks(randomIntSlice), time.Since(t).Milliseconds())
}
