package sorting

import . "gopkg.in/check.v1"

func (s *SortingTestSuite) TestBubbleSort(c *C) {
	var b BubbleSort
	r := b.Sort(GetUnsortedList())

	c.Assert(r, DeepEquals, sorted_list)
	c.Assert(r, Not(DeepEquals), unsorted_list)
}

func (s *SortingTestSuite) BenchmarkBubbleSort(c *C) {
	list := GetUnsortedList()

	for i := 0; i < c.N; i++ {
		var b BubbleSort
		b.Sort(list)
	}
}
