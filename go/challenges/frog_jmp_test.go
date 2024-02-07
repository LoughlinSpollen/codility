package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestFrogJump(t *testing.T) {
	r := challenges.SolutionFrogJump(10, 85, 30)
	assert.Equal(t, 3, r, "Expected 3, but got %v", r)

	r = challenges.SolutionFrogJump(0, 1000000000, 1)
	assert.Equal(t, 1000000000, r, "Expected 1000000000, but got %v", r)

	r = challenges.SolutionFrogJump(0, 10, 20)
	assert.Equal(t, 1, r, "Expected 1, but got %v", r)

	r = challenges.SolutionFrogJump(10, 100, 10)
	assert.Equal(t, 9, r, "Expected 9, but got %v", r)

	r = challenges.SolutionFrogJump(10, 10, 10)
	assert.Equal(t, 0, r, "Expected 0, but got %v", r)

	r = challenges.SolutionFrogJump(9, 29, 10)
	assert.Equal(t, 2, r, "Expected 2, but got %v", r)

	r = challenges.SolutionFrogJump(1, 1, 1)
	assert.Equal(t, 0, r, "Expected 0, but got %v", r)

	r = challenges.SolutionFrogJump(1, 2, 1)
	assert.Equal(t, 1, r, "Expected 1, but got %v", r)

	r = challenges.SolutionFrogJump(1, 2, 2)
	assert.Equal(t, 1, r, "Expected 1, but got %v", r)

	r = challenges.SolutionFrogJump(1, 4, 1)
	assert.Equal(t, 3, r, "Expected 3, but got %v", r)

	r = challenges.SolutionFrogJump(1, 6, 1)
	assert.Equal(t, 5, r, "Expected 5, but got %v", r)
}
