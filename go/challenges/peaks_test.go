package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionPeaks(t *testing.T) {
	r := challenges.SolutionPeaks([]int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2})
	assert.Equal(t, 3, r, fmt.Sprintf("Expected 3, but got %v", r))

	r = challenges.SolutionPeaks([]int{1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1})
	assert.Equal(t, 10, r, fmt.Sprintf("Expected 10, but got %v", r))

	r = challenges.SolutionPeaks([]int{1, 2, 3, 4, 6, 5})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))
}
