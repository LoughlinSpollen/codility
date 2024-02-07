package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionEquiLeader(t *testing.T) {
	r := challenges.SolutionEquiLeader([]int{4, 3, 4, 4, 4, 2})
	assert.Equal(t, 2, r, fmt.Sprintf("Expected 2, but got %v", r))

	r = challenges.SolutionEquiLeader([]int{1})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionEquiLeader([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = challenges.SolutionEquiLeader([]int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	assert.Equal(t, 8, r, fmt.Sprintf("Expected 8, but got %v", r))
}
