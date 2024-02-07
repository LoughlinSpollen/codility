package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestFrogRiverOne(t *testing.T) {
	r := challenges.SolutionFrogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 5, 4})
	assert.Equal(t, 6, r, "Expected 6, but got %v", r)

	r = challenges.SolutionFrogRiverOne(1, []int{1})
	assert.Equal(t, 0, r, "Expected 0, but got %v", r)

	r = challenges.SolutionFrogRiverOne(2, []int{2, 1})
	assert.Equal(t, 1, r, "Expected 1, but got %v", r)

	r = challenges.SolutionFrogRiverOne(2, []int{1, 2})
	assert.Equal(t, 1, r, "Expected 1, but got %v", r)

	r = challenges.SolutionFrogRiverOne(2, []int{2})
	assert.Equal(t, -1, r, "Expected -1, but got %v", r)

	r = challenges.SolutionFrogRiverOne(2, []int{1, 1})
	assert.Equal(t, -1, r, "Expected -1, but got %v", r)

	r = challenges.SolutionFrogRiverOne(3, []int{1, 2})
	assert.Equal(t, -1, r, "Expected -1, but got %v", r)
}
