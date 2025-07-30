package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var wg sync.WaitGroup

type results struct{
	mu sync.Mutex
	resultsArray []int
}

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	list := make([]int, 0, size)
	if size < 1 {
		return []int {}
	}
	for range size {
		list = append(list, rand.Intn(size))
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
	var resultMaxChunks results
	if len(data) < CHUNKS^CHUNKS {
		return maximum(data)
	}
	sizeOfChunk := int(math.Ceil(float64(len(data)) / float64(CHUNKS)))
	for i := 1; i <= CHUNKS; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i == CHUNKS {
				chunk := data[(i - 1) * sizeOfChunk:]
			} else {
				chunk := data[(i - 1) * sizeOfChunk:i * sizeOfChunk - 1]
			}
			max := maximum(chunk)
			resultMaxChunks.mu.Lock()
			resultMaxChunks.resultsArray = append(resultMaxChunks.resultsArray, max)
			resultMaxChunks.mu.Unlock()
			return
		}(i)
	}
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь

	//fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь

	//fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
