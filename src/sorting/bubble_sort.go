package sorting

// BubbleSort is the slowest of all sorting algorithms
// with a running time of O(n^2).
type BubbleSort struct{}

func (b *BubbleSort) Sort(sorted []int) []int {
	size := len(sorted)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}

			// swap
			if sorted[i] < sorted[j] {
				sorted[j], sorted[i] = sorted[i], sorted[j]
			}
		}
	}
	return sorted
}
