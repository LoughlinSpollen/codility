package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestMaxNonoverlappingSegments(t *testing.T) {
	r := challenges.SolutionMaxNonoverlappingSegments([]int{1, 3, 7, 9, 9}, []int{5, 6, 8, 9, 10})
	assert.Equal(t, 3, r, "Expected 3, but got %v", r)
}
