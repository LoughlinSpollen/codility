package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestMaxProductOfThree(t *testing.T) {
	r := challenges.SolutionMaxProductOfThree([]int{-3, 1, 2, -2, 5, 6})
	assert.Equal(t, 60, r, "Expected 60, but got %v", r)

	r = challenges.SolutionMaxProductOfThree([]int{-3, -1, -2, -2, -5, -6})
	assert.Equal(t, -4, r, "Expected -4, but got %v", r)
}
