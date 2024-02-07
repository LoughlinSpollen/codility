package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestMinAvgTwoSlice(t *testing.T) {
	r := challenges.SolutionMinAvgTwoSlice([]int{4, 2, 2, 5, 1, 5, 8})
	assert.Equal(t, 1, r, "Expected 1, but got %v", r)

	r = challenges.SolutionMinAvgTwoSlice([]int{1, 1})
	assert.Equal(t, 0, r, "Expected 0, but got %v", r)

	r = challenges.SolutionMinAvgTwoSlice([]int{1, 1, 1, 1, 1, 1, 1})
	assert.Equal(t, 0, r, "Expected 0, but got %v", r)

	r = challenges.SolutionMinAvgTwoSlice([]int{5, 2, 2, 100, 1, 1, 100})
	assert.Equal(t, 4, r, "Expected 4, but got %v", r)

	r = challenges.SolutionMinAvgTwoSlice([]int{11, 2, 10, 1, 100, 2, 9, 2, 100})
	assert.Equal(t, 1, r, "Expected 1, but got %v", r)

	r = challenges.SolutionMinAvgTwoSlice([]int{1, -1, 1, -1})
	assert.Equal(t, 1, r, "Expected 1, but got %v", r)
}
