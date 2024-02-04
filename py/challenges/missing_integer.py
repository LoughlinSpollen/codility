import unittest
import random

# Write a function:

# def solution(A)

# that, given an array A of N integers, returns the smallest positive integer 
# (greater than 0) that does not occur in A.

# For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

# Given A = [1, 2, 3], the function should return 4.

# Given A = [−1, −3], the function should return 1.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [1..100,000];
# each element of array A is an integer within the range [−1,000,000..1,000,000].

def solution(A):
    intSet = set(A)

    i = 0
    missing = 1
    max = len(intSet)
    while(i < max):
        i += 1
        val = intSet.pop()                
        if val > 0: 
            if val != missing:
                return missing
            else: 
                missing += 1
            if i == max:
                return missing

    return missing



N_RANGE = (1, 100000)
ELEMENT_RANGE = (-2147483648, 2147483647)

class MissingIntegerTest(unittest.TestCase):
    def test_example(self):
        self.assertEqual(solution([1, 3, 6, 4, 1, 2]), 5)

    def test_simple(self):
        self.assertEqual(solution([1, 2, 3]), 4)

    def test_negative(self):
        self.assertEqual(solution([-1, -3]), 1)


    def test_extreme_single(self):
        self.assertEqual(solution([2]), 1)
        self.assertEqual(solution([1]), 2)
        self.assertEqual(solution([-1]), 1)

    def test_extreme_min_max_int(self):
        self.assertEqual(1, solution([ELEMENT_RANGE[0], ELEMENT_RANGE[1], -10]))

    def test_no_gaps(self):
        self.assertEqual(solution([1, 2, 3, 4, 5]), 6)

    def test_duplicates(self):
        self.assertEqual(solution([1, 1, 1, 3, 3, 3]), 2)

