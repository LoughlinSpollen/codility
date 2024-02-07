package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestMaxCounters(t *testing.T) {
	r := challenges.SolutionMaxCounters(5, []int{3, 4, 4, 6, 1, 4, 4})
	assert.Equal(t, []int{3, 2, 2, 4, 2}, r, "Expected [3, 2, 2, 4, 2], but got %v", r)
}
