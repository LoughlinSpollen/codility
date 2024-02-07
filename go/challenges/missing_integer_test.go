package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionMissingInteger(t *testing.T) {
	r := challenges.SolutionMissingInteger([]int{1, 3, 6, 4, 1, 2})
	assert.Equal(t, 5, r, fmt.Sprintf("Expected 5, but got %v", r))

	r = challenges.SolutionMissingInteger([]int{1, 2, 3})
	assert.Equal(t, 4, r, fmt.Sprintf("Expected 4, but got %v", r))

	r = challenges.SolutionMissingInteger([]int{-1, -3})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionMissingInteger([]int{2})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionMissingInteger([]int{1})
	assert.Equal(t, 2, r, fmt.Sprintf("Expected 2, but got %v", r))

	r = challenges.SolutionMissingInteger([]int{-1})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionMissingInteger([]int{2147483647, -2147483648, 0})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionMissingInteger([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 6, r, fmt.Sprintf("Expected 6, but got %v", r))

	r = challenges.SolutionMissingInteger([]int{1, 1, 1, 3, 3, 3})
	assert.Equal(t, 2, r, fmt.Sprintf("Expected 2, but got %v", r))
}
