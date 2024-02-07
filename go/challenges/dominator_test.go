package challenges_test

import (
	"fmt"
	"testing"

	. "github.com/loughlinspollen/codility/go/challenges"

	"github.com/stretchr/testify/assert"
)

func TestDominator(t *testing.T) {
	r := SolutionDominator([]int{3, 4, 3, 2, 3, -1, 3, 3})
	assert.Equal(t, 3, r, fmt.Sprintf("Expected 3, but got %v", r))

	r = SolutionDominator([]int{1})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = SolutionDominator([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = SolutionDominator([]int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = SolutionDominator([]int{1, 1, 1, 1, 1, 1, 1, 1, 2})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	var testArr []int
	for i := 0; i < 100000; i++ {
		testArr = append(testArr, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
	r = SolutionDominator(testArr)
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))
}
