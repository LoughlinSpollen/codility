package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionTriangle(t *testing.T) {
	r := challenges.SolutionTriangle([]int{10, 2, 5, 1, 8, 20})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionTriangle([]int{10, 50, 5, 1})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = challenges.SolutionTriangle([]int{1, 1, 1})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))
}
