package challenges

//  Located on a line are N segments, numbered from 0 to N − 1, whose positions are
//  given in arrays A and B. For each I (0 ≤ I < N) the position of segment I is
//  from A[I] to B[I] (inclusive). The segments are sorted by their ends, which means
//  that B[K] ≤ B[K + 1] for K such that 0 ≤ K < N − 1.

//  Two segments I and J, such that I ≠ J, are overlapping if they share at least
//  one common point. In other words, A[I] ≤ A[J] ≤ B[I] or A[J] ≤ A[I] ≤ B[J].

//  We say that the set of segments is non-overlapping if it contains no two overlapping
//  segments. The goal is to find the size of a non-overlapping set containing the
//  maximal number of segments.

//  For example, consider arrays A, B such that:

//      A[0] = 1    B[0] = 5
//      A[1] = 3    B[1] = 6
//      A[2] = 7    B[2] = 8
//      A[3] = 9    B[3] = 9
//      A[4] = 9    B[4] = 10
//  The segments are shown in the figure below.

//  The size of a non-overlapping set containing a maximal number of segments is 3.
//  For example, possible sets are {0, 2, 3}, {0, 2, 4}, {1, 2, 3} or {1, 2, 4}.
//  There is no non-overlapping set with four segments.

//  Write a function:

//  func solution(A []int, B []int) int

//  that, given two arrays A and B consisting of N integers, returns the size of a
//  non-overlapping set containing a maximal number of segments.

//  For example, given arrays A, B shown above, the function should return 3, as explained above.

//  Write an efficient algorithm for the following assumptions:

//  N is an integer within the range [0..30,000];
//  each element of arrays A and B is an integer within the range [0..1,000,000,000];
//  A[I] ≤ B[I], for each I (0 ≤ I < N);
//  B[K] ≤ B[K + 1], for each K (0 ≤ K < N − 1).

func SolutionMaxNonoverlappingSegments(A []int, B []int) int {
	startPosArr := A
	endPosArr := B

	count := 1
	sz := len(startPosArr)
	prevEndPos := endPosArr[0]
	for i := 1; i < sz; i++ {
		if prevEndPos < startPosArr[i] {
			count += 1
		}
		prevEndPos = endPosArr[i]
	}

	return count
}
