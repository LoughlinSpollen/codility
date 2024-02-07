package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestMinMaxDivision(t *testing.T) {
	r := challenges.SolutionMinMaxDivision(3, 5, []int{2, 1, 5, 1, 2, 2, 2})
	assert.Equal(t, 6, r, "Expected 6, but got %v", r)
}
