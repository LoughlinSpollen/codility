package challenges

import (
	"slices"
	"sort"
)

//  You are given two non-empty arrays A and B consisting of N integers. These arrays
//  represent N planks. More precisely, A[K] is the start and B[K] the end of the K−th plank.

//  Next, you are given a non-empty array C consisting of M integers. This array represents
//  M nails. More precisely, C[I] is the position where you can hammer in the I−th nail.

//  We say that a plank (A[K], B[K]) is nailed if there exists a nail C[I] such that A[K] ≤ C[I] ≤ B[K].

//  The goal is to find the minimum number of nails that must be used until all the planks are
//  nailed. In other words, you should find a value J such that all planks will be nailed after using
//  only the first J nails. More precisely, for every plank (A[K], B[K]) such that 0 ≤ K < N, there
//  should exist a nail C[I] such that I < J and A[K] ≤ C[I] ≤ B[K].

//  For example, given arrays A, B such that:

//      A[0] = 1    B[0] = 4
//      A[1] = 4    B[1] = 5
//      A[2] = 5    B[2] = 9
//      A[3] = 8    B[3] = 10
//  four planks are represented: [1, 4], [4, 5], [5, 9] and [8, 10].

//  Given array C such that:

//      C[0] = 4
//      C[1] = 6
//      C[2] = 7
//      C[3] = 10
//      C[4] = 2
//  if we use the following nails:

//  0, then planks [1, 4] and [4, 5] will both be nailed.
//  0, 1, then planks [1, 4], [4, 5] and [5, 9] will be nailed.
//  0, 1, 2, then planks [1, 4], [4, 5] and [5, 9] will be nailed.
//  0, 1, 2, 3, then all the planks will be nailed.
//  Thus, four is the minimum number of nails that, used sequentially, allow all the planks to be nailed.

//  Write a function:

//  func solution(A []int, B []int, C []int) int

//  that, given two non-empty arrays A and B consisting of N integers and a non-empty array C consisting
//  of M integers, returns the minimum number of nails that, used sequentially, allow all the planks to
//  be nailed.

//  If it is not possible to nail all the planks, the function should return −1.

//  For example, given arrays A, B, C such that:

//      A[0] = 1    B[0] = 4
//      A[1] = 4    B[1] = 5
//      A[2] = 5    B[2] = 9
//      A[3] = 8    B[3] = 10

//      C[0] = 4
//      C[1] = 6
//      C[2] = 7
//      C[3] = 10
//      C[4] = 2
//  the function should return 4, as explained above.

//  Write an efficient algorithm for the following assumptions:

//  N and M are integers within the range [1..30,000];
//  each element of arrays A, B and C is an integer within the range [1..2*M];
//  A[K] ≤ B[K].

func SolutionNailingPlanks(A []int, B []int, C []int) int {
	start := A
	end := B
	nailsArr := C

	if slices.Min(start) > slices.Max(nailsArr) || slices.Max(start) < slices.Min(nailsArr) {
		return 0
	}

	//  Create a list of tuples with the start and end of the planks
	planks := make([][]int, len(start))
	for i, _ := range planks {
		planks[i] = []int{start[i], end[i]}
	}
	sort.Slice(planks, func(i, j int) bool {
		return planks[i][1] < planks[j][1]
	})

	binarySearchPlank := func(nail int) int {
		lowerBound := 0
		upperBound := len(planks) - 1
		for lowerBound <= upperBound {
			midPt := int((lowerBound + upperBound) / 2)
			if planks[midPt][1] < nail {
				lowerBound = midPt + 1
			} else if planks[midPt][0] > nail {
				upperBound = midPt - 1
			} else { //  plank found
				return midPt
			}
		}
		return -1
	}

	for count, nail := range nailsArr {
		i := binarySearchPlank(nail)
		for i > -1 {
			planks = append(planks[:i], planks[i+1:]...)
			i = binarySearchPlank(nail)
		}
		if len(planks) == 0 {
			return count + 1
		}
	}

	return -1
}
