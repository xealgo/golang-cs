package sorting

import . "gopkg.in/check.v1"

func (s *SortingTestSuite) TestMergeSortMerge(c *C) {
	var m MergeSort

	a := []int{2, 3}
	b := []int{1, 4}
	e := []int{1, 2, 3, 4}

	r := m.merge(a, b)

	c.Assert(r, DeepEquals, e)
}

func (s *SortingTestSuite) TestMergeSortDivide(c *C) {
	var m MergeSort

	l := []int{2, 5, 1, 4, 3, 6}
	e1 := []int{2, 5, 1}
	e2 := []int{4, 3, 6}

	a, b := m.divide(l)

	c.Assert(a, DeepEquals, e1)
	c.Assert(b, DeepEquals, e2)
}

func (s *SortingTestSuite) TestMergeSortDivideAndMerge(c *C) {
	var m MergeSort
	l := []int{14, 7, 3, 12, 9, 11, 6, 2}
	e := []int{2, 3, 6, 7, 9, 11, 12, 14}

	// first split l1, l2
	l1, r1 := m.divide(l)
	c.Assert(l1, DeepEquals, []int{14, 7, 3, 12})
	c.Assert(r1, DeepEquals, []int{9, 11, 6, 2})

	l2, r2 := m.divide(l1)
	l3, r3 := m.divide(r1)
	l4, r4 := m.divide(l2)
	l5, r5 := m.divide(r2)
	l6, r6 := m.divide(l3)
	l7, r7 := m.divide(r3)

	c.Assert(l7, DeepEquals, []int{6})
	c.Assert(r7, DeepEquals, []int{2})

	m1 := m.merge(l4, r4)
	m2 := m.merge(l5, r5)
	m3 := m.merge(l6, r6)
	m4 := m.merge(l7, r7)

	c.Assert(m3, DeepEquals, []int{9, 11})
	c.Assert(m4, DeepEquals, []int{2, 6})

	m5 := m.merge(m1, m2)
	m6 := m.merge(m3, m4)
	m7 := m.merge(m5, m6)

	c.Assert(m7, DeepEquals, e)
}

func (s *SortingTestSuite) TestMergeSortSimple(c *C) {
	var m MergeSort

	l := []int{14, 7, 3, 12, 9, 11, 6, 2}
	e := []int{2, 3, 6, 7, 9, 11, 12, 14}
	r := m.Sort(l)

	c.Assert(r, DeepEquals, e)
}

func (s *SortingTestSuite) TestMergeSort(c *C) {
	var m MergeSort
	r := m.Sort(GetUnsortedList())

	c.Assert(r, DeepEquals, sorted_list)
}

func (s *SortingTestSuite) BenchmarkMergeSort(c *C) {
	list := GetUnsortedList()

	for i := 0; i < c.N; i++ {
		var m MergeSort
		m.Sort(list)
	}
}
