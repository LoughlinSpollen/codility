package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionCountFactors(t *testing.T) {
	r := challenges.SolutionCountFactors(24)
	assert.Equal(t, 8, r, fmt.Sprintf("Expected 8, but got %v", r))

	r = challenges.SolutionCountFactors(1)
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))
}
