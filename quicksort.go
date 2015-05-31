package algorithms

func partition(A []int, p, r int) int {
	x := A[r] // Pivot element
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			i = i + 1
			swap(A, i, j)
		}
	}
	swap(A, i + 1, r)
	return i + 1
}

func quickSort(A []int, p, r int) []int {
	if (p < r) {
		q := partition(A, p, r)
		quickSort(A, p, q - 1)
		quickSort(A, q + 1, r)
	}
	return A
}

func QuickSort(unsorted []int) []int {
	return quickSort(unsorted, 0, len(unsorted) - 1)
}
