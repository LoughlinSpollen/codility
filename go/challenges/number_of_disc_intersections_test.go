package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionNumberOfDiscIntersections(t *testing.T) {
	r := challenges.SolutionNumberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0})
	assert.Equal(t, 11, r, fmt.Sprintf("Expected 11, but got %v", r))

	r = challenges.SolutionNumberOfDiscIntersections([]int{3, 2, 1})
	assert.Equal(t, 3, r, fmt.Sprintf("Expected 3, but got %v", r))

	r = challenges.SolutionNumberOfDiscIntersections([]int{1, 1, 1})
	assert.Equal(t, 3, r, fmt.Sprintf("Expected 3, but got %v", r))

	r = challenges.SolutionNumberOfDiscIntersections([]int{})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))
}
