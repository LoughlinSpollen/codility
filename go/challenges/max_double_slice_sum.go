package challenges

// A non-empty array A consisting of N integers is given.

// A triplet (X, Y, Z), such that 0 ≤ X < Y < Z < N, is called a double slice.

// The sum of double slice (X, Y, Z) is the total of
// A[X + 1] + A[X + 2] + ... + A[Y − 1] + A[Y + 1] + A[Y + 2] + ... + A[Z − 1].

// For example, array A such that:

//     A[0] = 3
//     A[1] = 2
//     A[2] = 6
//     A[3] = -1
//     A[4] = 4
//     A[5] = 5
//     A[6] = -1
//     A[7] = 2
// contains the following example double slices:

// double slice (0, 3, 6), sum is 2 + 6 + 4 + 5 = 17,
// double slice (0, 3, 7), sum is 2 + 6 + 4 + 5 − 1 = 16,
// double slice (3, 4, 5), sum is 0.
// The goal is to find the maximal sum of any double slice.

// Write a function:

// func solution(A []int) int

// that, given a non-empty array A consisting of N integers, returns the maximal sum of any double slice.

// For example, given:

//     A[0] = 3
//     A[1] = 2
//     A[2] = 6
//     A[3] = -1
//     A[4] = 4
//     A[5] = 5
//     A[6] = -1
//     A[7] = 2
// the function should return 17, because no double slice of array A has a sum of greater than 17.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [3..100,000];
// each element of array A is an integer within the range [−10,000..10,000].

func SolutionMaxDoubleSliceSum(A []int) int {
	intArr := A
	size := len(intArr)

	if size == 3 {
		return 0
	}

	kadaneMaxSlice := func(start, end, step int) []int {
		currentMaxVal := max(intArr[start], 0)
		maxSumArray := []int{0, currentMaxVal}
		for i := start + step; i != end; i += step {
			currentMaxVal = max(currentMaxVal, 0) + intArr[i]
			maxSumArray = append(maxSumArray, currentMaxVal)
		}
		return maxSumArray
	}

	maxLeftSlices := kadaneMaxSlice(1, size-1, 1)
	maxRightSlices := kadaneMaxSlice(size-2, 0, -1)

	maxSliceSum := 0
	size = len(maxLeftSlices)
	for i := 1; i < size-1; i++ {
		l := maxLeftSlices[i]
		r := maxRightSlices[size-i-2]
		maxSliceSum = max(maxSliceSum, l+r)
	}

	return maxSliceSum
}
