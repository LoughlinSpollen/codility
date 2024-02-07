package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionNesting(t *testing.T) {
	r := challenges.SolutionNesting("(()(())())")
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionNesting("())")
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = challenges.SolutionNesting("")
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))

	r = challenges.SolutionNesting(")")
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	r = challenges.SolutionNesting("(")
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	max := 100000
	strArr := make([]rune, max)
	for i := 0; i < max; i++ {
		strArr[i] = '('
	}
	r = challenges.SolutionNesting(string(strArr))
	assert.Equal(t, 0, r, fmt.Sprintf("Expected 0, but got %v", r))

	strArr1 := make([]rune, max)
	for i := 0; i < max/2; i++ {
		strArr1[i] = '('
	}
	for i := max / 2; i < max; i++ {
		strArr[i] = ')'
	}
	r = challenges.SolutionNesting(string(strArr1))
	assert.Equal(t, 1, r, fmt.Sprintf("Expected 1, but got %v", r))
}
