package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGenerateRandomElements(t *testing.T) {
	tests := []int{2, 100, 0, 232423}
	for _, i := range tests {
		assert.Len(t, generateRandomElements(i), i)
	}
}

func TestMaximum(t *testing.T) {
	type TestMax struct {
		data []int
		expected int
	}
	tests := []TestMax{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{}, 0},
		{[]int{0, 0, 0, 0}, 0},
		{[]int{122, 122, 5, 40}, 122},
		{[]int{2322, 5, 5, 47}, 2322},
	}
	for _, v := range tests {
		assert.Equal(t, maximum(v.data), v.expected)
	}
}