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