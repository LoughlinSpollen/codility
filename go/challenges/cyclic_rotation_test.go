package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestCyclicRotation(t *testing.T) {
	r := challenges.SolutionCyclicRotation([]int{6, 3, 8, 9, 7}, 0)
	assert.Equal(t, []int{6, 3, 8, 9, 7}, r, fmt.Sprintf("Expected [6, 3, 8, 9, 7], but got %v", r))

	r = challenges.SolutionCyclicRotation([]int{3, 8, 9, 7, 6}, 3)
	assert.Equal(t, []int{9, 7, 6, 3, 8}, r, fmt.Sprintf("Expected [9, 7, 6, 3, 8], but got %v", r))

	r = challenges.SolutionCyclicRotation([]int{}, 5)
	assert.Equal(t, []int{}, r, fmt.Sprintf("Expected [], but got %v", r))

	r = challenges.SolutionCyclicRotation([]int{1, 2, 3, 4}, 1)
	assert.Equal(t, []int{4, 1, 2, 3}, r, fmt.Sprintf("Expected [4, 1, 2, 3], but got %v", r))

	r = challenges.SolutionCyclicRotation([]int{1, 2, 3, 4}, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, r, fmt.Sprintf("Expected [1, 2, 3, 4], but got %v", r))
}
