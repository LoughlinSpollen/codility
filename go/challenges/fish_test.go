package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionFish(t *testing.T) {
	r := challenges.SolutionFish([]int{4, 3, 2, 1, 5}, []int{0, 1, 0, 0, 0})
	assert.Equal(t, 2, r, fmt.Sprintf("Expected 2, but got %v", r))
}
