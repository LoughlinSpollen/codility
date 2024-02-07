package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	r := challenges.SolutionMaxProfit([]int{23171, 21011, 21123, 21366, 21013, 21367})
	assert.Equal(t, 356, r, "Expected 356, but got %v", r)

	r = challenges.SolutionMaxProfit([]int{1, 2})
	assert.Equal(t, 1, r, "Expected 1, but got %v", r)

	r = challenges.SolutionMaxProfit([]int{-1, -2})
	assert.Equal(t, 0, r, "Expected 0, but got %v", r)

	r = challenges.SolutionMaxProfit([]int{})
	assert.Equal(t, 0, r, "Expected 0, but got %v", r)

	r = challenges.SolutionMaxProfit([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, 9, r, "Expected 9, but got %v", r)

	r = challenges.SolutionMaxProfit([]int{200000, 200000, 200000, 200000, 200000, 200000})
	assert.Equal(t, 0, r, "Expected 0, but got %v", r)
}
