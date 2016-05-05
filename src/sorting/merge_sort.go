package sorting

import "math"

type MergeSort struct{}

// Sort the primary sort function
func (m *MergeSort) Sort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	// initial divide
	l1, r1 := m.divide(list)
	a := m.Sort(l1)
	b := m.Sort(r1)

	list = m.merge(a, b)
	return list
}

// divide splits an array into 2 "equal" parts
func (m *MergeSort) divide(list []int) ([]int, []int) {
	q := int(math.Floor(float64(len(list)) / 2))
	a := list[0:q]
	b := list[q:]
	return a, b
}

// merge quickly sorts & merges 2 arrays
func (m *MergeSort) merge(a []int, b []int) []int {
	merged := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	l1, l2 := len(a), len(b)

	for i < l1 && j < l2 {
		if a[i] < b[j] {
			merged[k] = a[i]
			i++

		} else {
			merged[k] = b[j]
			j++
		}
		k++
	}

	for i < l1 {
		merged[k] = a[i]
		k++
		i++
	}

	for j < l2 {
		merged[k] = b[j]
		k++
		j++
	}

	return merged
}
