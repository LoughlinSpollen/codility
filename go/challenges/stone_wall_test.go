package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionStoneWall(t *testing.T) {
	r := challenges.SolutionStoneWall([]int{8, 8, 5, 7, 9, 8, 7, 4, 8})
	assert.Equal(t, 7, r, fmt.Sprintf("Expected 7, but got %v", r))

	r = challenges.SolutionStoneWall([]int{1})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionStoneWall([]int{1, 1})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionStoneWall([]int{1, 1, 1})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionStoneWall([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	assert.Equal(t, 9, r, fmt.Sprintf("Expected 9, but got %v", r))

	r = challenges.SolutionStoneWall([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	assert.Equal(t, 9, r, fmt.Sprintf("Expected 9, but got %v", r))
}
