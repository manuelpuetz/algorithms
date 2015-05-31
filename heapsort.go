package heapsort

type Heap struct {
	Field      []int
	HeapLength int
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2 * i + 1
}

func right(i int) int {
	return 2 * i + 2
}

func swap(array []int, i, j int) []int {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
	return array
}

func maxHeapify(A Heap, i int) Heap {
	var maximum int
	l := left(i)
	r := right(i)
	if l < A.HeapLength && A.Field[l] > A.Field[i] {
		maximum = l
	} else {
		maximum = i
	}
	if r < A.HeapLength && A.Field[r] > A.Field[maximum] {
		maximum = r
	}

	if maximum != i {
		swap(A.Field, i, maximum)
		maxHeapify(A, maximum)
	}
	return A
}

func buildMaxHeap(A *Heap) Heap {
	(*A).HeapLength = len((*A).Field)
	for i := len((*A).Field)/2; i >= 0; i-- {
		maxHeapify(*A, i)
	}
	return *A
}

func Sort(unsorted []int) []int {
	A := Heap{Field: unsorted}
	buildMaxHeap(&A)
	for i := len(A.Field) - 1; i >= 1; i-- {
		swap(A.Field, 0, i)
		A.HeapLength = A.HeapLength - 1
		maxHeapify(A, 0)
	}
	return A.Field
}
