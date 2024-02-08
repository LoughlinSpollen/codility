package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionTapeEquilibrium(t *testing.T) {
	r := challenges.SolutionTapeEquilibrium([]int{3, 1, 2, 4, 3})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionTapeEquilibrium([]int{1, 2})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionTapeEquilibrium([]int{-1000, 1000})
	assert.Equal(t, 2000, r, fmt.Sprintf("Expected 2000, but got %v", r))
}
