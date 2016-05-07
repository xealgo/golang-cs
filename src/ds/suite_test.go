package ds

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

// Data Structures test suite
type DSTestSuite struct{}

// Registers the ds test suite
var _ = Suite(&DSTestSuite{})
