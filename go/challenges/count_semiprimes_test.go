package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestCountSemiprimes(t *testing.T) {
	r := challenges.SolutionCountSemiprimes(26, []int{1, 4, 16}, []int{26, 10, 20})
	assert.Equal(t, []int{10, 4, 0}, r, fmt.Sprintf("Expected [10, 4, 0], but got %v", r))

	r = challenges.SolutionCountSemiprimes(1, []int{1}, []int{1})
	assert.Equal(t, []int{0}, r, fmt.Sprintf("Expected [0], but got %v", r))

	r = challenges.SolutionCountSemiprimes(50000, []int{1}, []int{50000})
	assert.Equal(t, []int{12110}, r, fmt.Sprintf("Expected [12110], but got %v", r))
}
