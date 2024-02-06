package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionBinaryGap(t *testing.T) {
	N := 1041
	r := challenges.SolutionBinaryGap(N)
	assert.Equal(t, 5, r, fmt.Sprintf("Expected 5, but got %v", r))

	N = 32
	r = challenges.SolutionBinaryGap(N)
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	N = 1041
	r = challenges.SolutionBinaryGap(N)
	assert.Equal(t, 5, r, fmt.Sprintf("Expected 5, but got %v", r))
}
