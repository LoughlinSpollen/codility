package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestMaxSliceSum(t *testing.T) {
	r := challenges.SolutionMaxSliceSum([]int{3, 2, -6, 4, 0})
	assert.Equal(t, 5, r, "Expected 5, but got %v", r)

	r = challenges.SolutionMaxSliceSum([]int{-1, -1, -1, -1, -1})
	assert.Equal(t, -1, r, "Expected -1, but got %v", r)

	r = challenges.SolutionMaxSliceSum([]int{1, 1, 1, 1, 1})
	assert.Equal(t, 5, r, "Expected 5, but got %v", r)
}
