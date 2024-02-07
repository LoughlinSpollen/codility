package challenges_test

import (
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestMinPerimeterRectangle(t *testing.T) {
	r := challenges.SolutionMinPerimeterRectangle(30)
	assert.Equal(t, 22, r, "Expected 22, but got %v", r)

	r = challenges.SolutionMinPerimeterRectangle(1)
	assert.Equal(t, 4, r, "Expected 4, but got %v", r)
}
