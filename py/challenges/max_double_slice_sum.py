
import unittest

# A non-empty array A consisting of N integers is given.

# A triplet (X, Y, Z), such that 0 ≤ X < Y < Z < N, is called a double slice.

# The sum of double slice (X, Y, Z) is the total of 
# A[X + 1] + A[X + 2] + ... + A[Y − 1] + A[Y + 1] + A[Y + 2] + ... + A[Z − 1].

# For example, array A such that:

#     A[0] = 3
#     A[1] = 2
#     A[2] = 6
#     A[3] = -1
#     A[4] = 4
#     A[5] = 5
#     A[6] = -1
#     A[7] = 2
# contains the following example double slices:

# double slice (0, 3, 6), sum is 2 + 6 + 4 + 5 = 17,
# double slice (0, 3, 7), sum is 2 + 6 + 4 + 5 − 1 = 16,
# double slice (3, 4, 5), sum is 0.
# The goal is to find the maximal sum of any double slice.

# Write a function:

# def solution(A)

# that, given a non-empty array A consisting of N integers, returns the maximal sum of any double slice.

# For example, given:

#     A[0] = 3
#     A[1] = 2
#     A[2] = 6
#     A[3] = -1
#     A[4] = 4
#     A[5] = 5
#     A[6] = -1
#     A[7] = 2
# the function should return 17, because no double slice of array A has a sum of greater than 17.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [3..100,000];
# each element of array A is an integer within the range [−10,000..10,000].



N_RANGE = [3, 100000]
EL_RANGE = [-10000, 10000]

def solution(A):
    intArr = A
    size = len(intArr)

    if size == 3:
        return 0
    
    def kadaneMaxSlice(start, end, step=1):
        currentMaxVal = max(intArr[start], 0)
        maxSumArray = [0, currentMaxVal]
        for i in range(start+step, end, step):
            currentMaxVal = max(currentMaxVal, 0) + intArr[i]
            maxSumArray.append(currentMaxVal)
        return maxSumArray   


    maxLeftSlices = kadaneMaxSlice(1, size-1, 1)
    maxRightSlices = kadaneMaxSlice(size-2, 0, -1) 

    maxSliceSum = 0
    size = len(maxLeftSlices)
    for i in range(1, size-1):
        l = maxLeftSlices[i]
        r = maxRightSlices[size - i - 2]
        maxSliceSum = max(maxSliceSum, l + r)
        
    return maxSliceSum



class TestMaxDoubleSliceSum(unittest.TestCase):
    def test_ex(self):
        self.assertEqual(solution([3, 2, 6, -1, 4, 5, -1, 2]), 17)

