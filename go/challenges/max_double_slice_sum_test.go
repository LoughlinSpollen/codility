package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestMaxDoubleSliceSum(t *testing.T) {
	r := challenges.SolutionMaxDoubleSliceSum([]int{3, 2, 6, -1, 4, 5, -1, 2})
	assert.Equal(t, 17, r, "Expected 17, but got %v", r)
}
