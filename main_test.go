package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGenerateRandomElements(t *testing.T) {
	testsPositive := []int{2, 100, 0, 232423}
	testsNegative := []int{-2354, -934578, -0}
	for _, i := range testsPositive {
		assert.Len(t, generateRandomElements(i), i)
	}
	for _, i := range testsNegative {
		assert.Len(t, generateRandomElements(i), 0)
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