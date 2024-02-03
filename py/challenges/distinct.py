import unittest
import random 

# Write a function

# def solution(A)

# that, given an array A consisting of N integers, returns the number of distinct values in array A.

# For example, given array A consisting of six elements such that:

#  A[0] = 2    A[1] = 1    A[2] = 1
#  A[3] = 2    A[4] = 3    A[5] = 1
# the function should return 3, because there are 3 distinct values appearing in array A, namely 1, 2 and 3.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [0..100,000];


def solution(A):
        return len(set(A))

RANGE_A = (-1000000, 1000000)
RANGE_N = (0, 100000)
class DistinctTest(unittest.TestCase):
    def test_example(self):
        self.assertEqual(solution([2, 1, 1, 2, 3, 1]), 3)

    def test_simple(self):
        self.assertEqual(solution([0, 1, 2, 3, 4]), 5)

    def test_edges(self):
        self.assertEqual(solution([]), 0)
        self.assertEqual(solution([0]), 1)
        self.assertEqual(solution([1]), 1)
        self.assertEqual(solution([0, 1]), 2)
        self.assertEqual(solution([-1, 1]), 2)
        self.assertEqual(solution([RANGE_A[0], RANGE_A[1]]), 2)

    def test_medium(self):
        self.assertEqual(solution([1] * 500), 1)
        self.assertEqual(solution([n for n in range(-250, 250)]), 500)
        self.assertEqual(solution([n for n in range(-500, 500, 2)]), 500)
        self.assertEqual(solution([n for n in range(-500, 500, 2)] * 2), 500)

    def test_extreme(self):
        A = [n for n in range(*RANGE_N)]
        self.assertEqual(solution(A), RANGE_N[1])