package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestGenomicRangeQuery(t *testing.T) {
	r := challenges.SolutionGenomicRangeQuery("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6})
	assert.Equal(t, []int{2, 4, 1}, r, "Expected [2, 4, 1], but got %v", r)
}
