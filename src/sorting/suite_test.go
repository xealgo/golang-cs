package sorting

import (
	"testing"

	. "gopkg.in/check.v1"
)

// test data
var (
	sorted_list   = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
	unsorted_list = []int{29, 52, 81, 74, 22, 1, 76, 16, 66, 44, 97, 47, 53, 79, 62, 64, 71, 2, 87, 39, 38, 63, 9, 5, 34, 85, 0, 83, 41, 48, 8, 84, 7, 37, 15, 68, 51, 89, 40, 14, 19, 73, 36, 12, 96, 26, 3, 98, 20, 60, 24, 46, 4, 31, 80, 82, 33, 88, 86, 90, 6, 59, 91, 43, 61, 50, 28, 35, 72, 30, 42, 55, 18, 21, 17, 32, 70, 78, 69, 45, 49, 27, 65, 10, 58, 23, 57, 25, 56, 67, 94, 95, 77, 92, 54, 11, 75, 99, 13, 93}
)

func GetUnsortedList() []int {
	list := make([]int, len(unsorted_list))
	copy(list, unsorted_list)
	return list
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

// Sorting test suite
type SortingTestSuite struct{}

// Registers the sorting test suite
var _ = Suite(&SortingTestSuite{})
