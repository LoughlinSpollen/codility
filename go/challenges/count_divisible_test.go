package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionCountDivisible(t *testing.T) {
	r := challenges.SolutionCountDivisible(6, 11, 2)
	assert.Equal(t, 3, r, fmt.Sprintf("Expected 3, but got %v", r))

	r = challenges.SolutionCountDivisible(0, 0, 11)
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionCountDivisible(1, 1, 11)
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = challenges.SolutionCountDivisible(0, 2000000000, 1)
	assert.Equal(t, 2000000001, r, fmt.Sprintf("Expected 2000000001, but got %v", r))

	r = challenges.SolutionCountDivisible(0, 2000000000, 2000000000)
	assert.Equal(t, 2, r, fmt.Sprintf("Expected 2, but got %v", r))

	r = challenges.SolutionCountDivisible(0, 2000000000, 1999999999)
	assert.Equal(t, 2, r, fmt.Sprintf("Expected 2, but got %v", r))
}
