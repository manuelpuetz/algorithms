package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}, QuickSort([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}))
}

func TestPartition(t *testing.T) {
	slice := []int{1, 7, 3}
	assert.Equal(t, 1, partition(slice, 0, 2))
	assert.Equal(t, []int{1, 3, 7}, slice)
}