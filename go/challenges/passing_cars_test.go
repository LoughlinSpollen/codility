package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionPassingCars(t *testing.T) {
	r := challenges.SolutionPassingCars([]int{0, 1, 0, 1, 1})
	assert.Equal(t, 5, r, fmt.Sprintf("Expected 5, but got %v", r))
}
