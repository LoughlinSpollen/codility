package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionSpaceCrew(t *testing.T) {
	T := []int{6, 4, 7}
	D := []int{1, 3, 4}
	r := challenges.SolutionSpaceCrews(T, D)
	assert.Equal(t, 840, r, fmt.Sprintf("Expected 840, but got %v", r))
}
