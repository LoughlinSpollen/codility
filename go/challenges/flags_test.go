package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestFlags(t *testing.T) {
	r := challenges.SolutionFlags([]int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2})
	assert.Equal(t, 3, r, fmt.Sprintf("Expected 3, but got %v", r))

	r = challenges.SolutionFlags([]int{1, 6, 3, 6, 5, 6, 5, 6, 3, 2, 6, 1})
	assert.Equal(t, 3, r, fmt.Sprintf("Expected 3, but got %v", r))

	r = challenges.SolutionFlags([]int{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0})
	assert.Equal(t, 4, r, fmt.Sprintf("Expected 4, but got %v", r))

	r = challenges.SolutionFlags([]int{0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0})
	assert.Equal(t, 7, r, fmt.Sprintf("Expected 7, but got %v", r))

	r = challenges.SolutionFlags([]int{5, 6, 5, 6, 1, 1, 1, 1, 6, 5, 6, 5})
	assert.Equal(t, 4, r, fmt.Sprintf("Expected 4, but got %v", r))

	r = challenges.SolutionFlags([]int{1, 6, 1, 1, 1, 1, 1, 6, 1, 1, 1, 1, 1, 1, 1, 1, 6, 1, 1, 1, 1, 1, 6, 1})
	assert.Equal(t, 4, r, fmt.Sprintf("Expected 4, but got %v", r))

	r = challenges.SolutionFlags([]int{1})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = challenges.SolutionFlags([]int{1, 2})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = challenges.SolutionFlags([]int{1, 2, 3})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = challenges.SolutionFlags([]int{1, 2, 3, 4})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = challenges.SolutionFlags([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = challenges.SolutionFlags([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))
}
