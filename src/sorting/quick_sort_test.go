package sorting

import . "gopkg.in/check.v1"

func (s *SortingTestSuite) TestQuickSort(c *C) {
	// var q QuickSort
	// r := q.Sort(GetUnsortedList())
}

func (s *SortingTestSuite) BenchmarkQuickSort(c *C) {
	list := GetUnsortedList()

	for i := 0; i < c.N; i++ {
		var q QuickSort
		q.Sort(list)
	}
}
