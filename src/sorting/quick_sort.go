package sorting

import "math"

type QuickSort struct{}

func (q *QuickSort) Sort(list []int) []int {
	size := len(list)
	sorted := make([]int, size)

	copy(sorted, list)

	if size <= 1 {
		return sorted
	}

	return sorted
}

func (q *QuickSort) pivot(size int) int {
	return int(math.Ceil(float64(size) / 2))
}

func (q *QuickSort) swap(list []int, a int, b int) {
	list[a], list[b] = list[b], list[a]
}
