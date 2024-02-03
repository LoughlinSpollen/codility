import unittest


# A non-empty array A consisting of N integers is given. A pair of integers (P, Q), 
# such that 0 ≤ P < Q < N, is called a slice of array A (notice that the slice 
# contains at least two elements). The average of a slice (P, Q) is the sum of 
# A[P] + A[P + 1] + ... + A[Q] divided by the length of the slice. To be precise, 
# the average equals (A[P] + A[P + 1] + ... + A[Q]) / (Q − P + 1).

# For example, array A such that:

#     A[0] = 4
#     A[1] = 2
#     A[2] = 2
#     A[3] = 5
#     A[4] = 1
#     A[5] = 5
#     A[6] = 8
# contains the following example slices:

# slice (1, 2), whose average is (2 + 2) / 2 = 2;
# slice (3, 4), whose average is (5 + 1) / 2 = 3;
# slice (1, 4), whose average is (2 + 2 + 5 + 1) / 4 = 2.5.
# The goal is to find the starting position of a slice whose average is minimal.

# Write a function:

# def solution(A)

# that, given a non-empty array A consisting of N integers, returns the starting 
# position of the slice with the minimal average. If there is more than one slice 
# with a minimal average, you should return the smallest starting position of such a slice.

# For example, given array A such that:

#     A[0] = 4
#     A[1] = 2
#     A[2] = 2
#     A[3] = 5
#     A[4] = 1
#     A[5] = 5
#     A[6] = 8
# the function should return 1, as explained above.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [2..100,000];
# each element of array A is an integer within the range [−10,000..10,000].
                                                        



def solution(A):
    intArr = A
    size = len(intArr)
            
    sumArr = [0] * (size+1)
    for i in range(1, size+1):
        sumArr[i] = sumArr[i-1] + intArr[i-1]

    startPos = 0
    avg = -1
    for end in range (2, len(sumArr)):
        for start in range (1, end):
            sliceAvg = float((sumArr[end] - sumArr[start-1])) / (end-start+1)
            if avg is -1 or sliceAvg < avg: 
                avg = sliceAvg
                startPos = start
    
    return startPos -1

class MinAvgTwoSliceTest(unittest.TestCase):
    def test_example(self):        
        self.assertEqual(solution([4, 2, 2, 5, 1, 5, 8]), 1)

    def test_min(self):
        self.assertEqual(solution([1, 1]), 0)

    def test_max(self):
        self.assertEqual(solution([1, 1, 1, 1, 1, 1, 1]), 0)

    def test_two(self):
        self.assertEqual(solution([5, 2, 2, 100, 1, 1, 100]), 4)
        self.assertEqual(solution([11, 2, 10, 1, 100, 2, 9, 2, 100]), 1)

    def test_three(self):
        self.assertEqual(solution([1, -1, 1, -1]), 1)


