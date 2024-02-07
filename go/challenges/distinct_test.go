package challenges_test

import (
	"fmt"
	"testing"

	. "github.com/loughlinspollen/codility/go/challenges"

	"github.com/stretchr/testify/assert"
)

func TestSolutionDistinct(t *testing.T) {

	var rangeA = [...]int{-1000000, 1000000}
	var rangeN = [...]int{0, 100000}

	r := SolutionDistinct([]int{2, 1, 1, 2, 3, 1})
	assert.Equal(t, 3, r, fmt.Sprintf("Expected 3, but got %v", r))

	r = SolutionDistinct([]int{0, 1, 2, 3, 4})
	assert.Equal(t, 5, r, fmt.Sprintf("Expected 5, but got %v", r))

	r = SolutionDistinct([]int{})
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = SolutionDistinct([]int{0})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = SolutionDistinct([]int{1})
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = SolutionDistinct([]int{0, 1})
	assert.Equal(t, 2, r, fmt.Sprintf("Expected 2, but got %v", r))

	r = SolutionDistinct([]int{-1, 1})
	assert.Equal(t, 2, r, fmt.Sprintf("Expected 2, but got %v", r))

	r = SolutionDistinct([]int{rangeA[0], rangeA[1]})
	assert.Equal(t, 2, r, fmt.Sprintf("Expected 2, but got %v", r))

	testArr := make([]int, 500)
	for i := range testArr {
		testArr[i] = int(1)
	}
	r = SolutionDistinct(testArr)
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	rangeArr := func(start, end int, step int) []int {
		var r []int
		for i := start; i < end; i += step {
			r = append(r, i)
		}
		return r
	}

	r = SolutionDistinct(rangeArr(-250, 250, 1))

	assert.Equal(t, 500, r, fmt.Sprintf("Expected 500, but got %v", r))

	r = SolutionDistinct(rangeArr(-500, 500, 2))
	assert.Equal(t, 500, r, fmt.Sprintf("Expected 500, but got %v", r))

	r = SolutionDistinct(rangeArr(rangeN[0], rangeN[1], 1))
	assert.Equal(t, rangeN[1], r, fmt.Sprintf("Expected %v, but got %v", rangeN[1], r))
}
