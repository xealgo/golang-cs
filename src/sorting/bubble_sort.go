package sorting

// BubbleSort is the slowest of all sorting algorithms
// with a running time of O(n^2).
type BubbleSort struct{}

func (b *BubbleSort) Sort(unsorted []int) []int {
	size := len(unsorted)
	sorted := make([]int, size)

	copy(sorted, unsorted)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}

			if sorted[i] < sorted[j] {
				var tmp = sorted[j]
				sorted[j] = sorted[i]
				sorted[i] = tmp
			}
		}
	}
	return sorted
}
