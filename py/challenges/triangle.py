import unittest
import random

# An array A consisting of N integers is given. A triplet (P, Q, R) is triangular if 0 ≤ P < Q < R < N and:

# A[P] + A[Q] > A[R],
# A[Q] + A[R] > A[P],
# A[R] + A[P] > A[Q].
# For example, consider array A such that:

#   A[0] = 10    A[1] = 2    A[2] = 5
#   A[3] = 1     A[4] = 8    A[5] = 20
# Triplet (0, 2, 4) is triangular.

# Write a function:

# def solution(A)

# that, given an array A consisting of N integers, returns 1 if there exists a triangular triplet for this array and returns 0 otherwise.

# For example, given array A such that:

#   A[0] = 10    A[1] = 2    A[2] = 5
#   A[3] = 1     A[4] = 8    A[5] = 20
# the function should return 1, as explained above. Given array A such that:

#   A[0] = 10    A[1] = 50    A[2] = 5
#   A[3] = 1
# the function should return 0.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [0..100,000];
# each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].

N_MAX = 100000
RANGE_A = (-2147483648, 2147483647)

def solution(A):
    intArr = A

    intArr.sort()
    max = len(intArr)
    if max >= 3:
        p = intArr[0]
        q = intArr[1]

        for i in range (2, max):
            r = intArr[i]
            if p >= 0 and p < q and q < r and r < N_MAX:
                if p + q > r and q + r > p and r + p > q:
                    return 1
            p = q
            q = r
    return 0



class TriangleTest(unittest.TestCase):
    def test_example(self):
        self.assertEqual(solution([10, 2, 5, 1, 8, 20]), 1)
        self.assertEqual(solution([10, 50, 5, 1]), 0)

    def test_extreme(self):
        arr = [random.randint(*RANGE_A)  for n in range(0, N_MAX)]
        res = solution(arr)
        print(res)

    def test_neg(self):
        arr = [-1] * N_MAX
        res = solution(arr)
        print(res)
    
    def test_empty(self):
        self.assertEqual(solution([]), 0)
        self.assertEqual(solution([1,1,1]), 0)

