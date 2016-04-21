package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCircleAreaComputation(c *C) {
	myCircle := circle{2}

	// Stop test execution if test fails.
	c.Assert(myCircle.radius, Equals, 2.0)

	// Continue test execution if test fails.
	c.Check(myCircle.radius, Equals, 2.0)
}
