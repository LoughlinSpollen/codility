package challenges_test

import (
	"fmt"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionBrackets(t *testing.T) {
	r := challenges.SolutionBrackets("{[()()]}")
	assert.Equal(t, true, r, fmt.Sprintf("Expected true, but got %v", r))

	r = challenges.SolutionBrackets("([)()]")
	assert.Equal(t, false, r, fmt.Sprintf("Expected false, but got %v", r))

	r = challenges.SolutionBrackets("(")
	assert.Equal(t, false, r, fmt.Sprintf("Expected false, but got %v", r))

	r = challenges.SolutionBrackets(")")
	assert.Equal(t, false, r, fmt.Sprintf("Expected false, but got %v", r))

	max := 3000
	brackets := ""
	closing := ""
	for i := 0; i < max; i++ {
		brackets += "("
		closing += ")"
	}
	r = challenges.SolutionBrackets(brackets + closing)
	assert.Equal(t, true, r, fmt.Sprintf("Expected true, but got %v", r))
}
