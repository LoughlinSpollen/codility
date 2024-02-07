package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionNailingPlanks(t *testing.T) {
	r := challenges.SolutionNailingPlanks([]int{1, 4, 5, 8}, []int{4, 5, 9, 10}, []int{4, 6, 7, 10, 2})
	assert.Equal(t, 4, r, fmt.Sprintf("Expected 4, but got %v", r))

	r = challenges.SolutionNailingPlanks([]int{1, 4}, []int{4, 5}, []int{4, 6, 7, 10, 2})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionNailingPlanks([]int{1, 4}, []int{4, 5}, []int{2, 6, 7, 10, 2})
	assert.Equal(t, -1, r, fmt.Sprintf("Expected -1, but got %v", r))

	r = challenges.SolutionNailingPlanks([]int{1, 1000000000}, []int{1000000000, 1000000000}, []int{1000000000})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionNailingPlanks([]int{1, 1000000000}, []int{1000000000, 1000000000}, []int{0})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))
}
