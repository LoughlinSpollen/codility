package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionOddOccurencesInArray(t *testing.T) {
	r := challenges.SolutionOddOccurencesInArray([]int{3, 9, 3, 9, 3, 3, 9, 7, 9})
	assert.Equal(t, 7, r, fmt.Sprintf("Expected 7, but got %v", r))
}
