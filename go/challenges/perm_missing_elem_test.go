package challenges_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/loughlinspollen/codility/go/challenges"
	"github.com/stretchr/testify/assert"
)

func TestSolutionPermMissingElem(t *testing.T) {
	r := challenges.SolutionPermMissingElem([]int{2, 3, 1, 5})
	assert.Equal(t, 4, r, fmt.Sprintf("Expected 4, but got %v", r))

	const max = 100000
	missing := rand.Intn(100)
	var intArr []int
	for i := 1; i <= max; i++ {
		if i == missing {
			continue
		}
		intArr = append(intArr, i)
	}

	r = challenges.SolutionPermMissingElem(intArr)
	assert.Equal(t, missing, r, fmt.Sprintf("Expected %v, but got %v", missing, r))
}
