package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}, HeapSort([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}))
}

func TestParent(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, parent(0))
	assert.Equal(0, parent(1))
	assert.Equal(0, parent(2))
	assert.Equal(3, parent(7))
	assert.Equal(3, parent(8))
}

func TestLeft(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, left(0))
	assert.Equal(7, left(3))
}

func TestRight(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, right(0))
	assert.Equal(10, right(4))
}

func TestSwap(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{3, 2, 1}, swap([]int{1, 2, 3}, 0, 2))
}

func TestMaxHeapify(t *testing.T) {
	assert.Equal(t, Heap{[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}, 10}, maxHeapify(Heap{[]int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}, 10}, 1))
}

func TestBuildMaxHeap(t *testing.T) {
	assert.Equal(t, Heap{[]int{1}, 1}, buildMaxHeap(&Heap{[]int{1}, 0}))
	assert.Equal(t, Heap{[]int{3, 1}, 2}, buildMaxHeap(&Heap{[]int{1, 3}, 0}))
}