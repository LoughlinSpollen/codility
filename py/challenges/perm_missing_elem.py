import unittest
import random

# An array A consisting of N different integers is given. The array contains integers in 
# the range [1..(N + 1)], which means that exactly one element is missing.

# Your goal is to find that missing element.

# Write a function:

# def solution(A)

# that, given an array A, returns the value of the missing element.

# For example, given array A such that:

#   A[0] = 2
#   A[1] = 3
#   A[2] = 1
#   A[3] = 5
# the function should return 4, as it is the missing element.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [0..100,000];
# the elements of A are all distinct;
# each element of array A is an integer within the range [1..(N + 1)].


def solution(A):

    intArr = A

    if len(intArr) == 0:
        return 1

    return sum(range(1, len(intArr)+2)) - sum(intArr)


INT_RANGE = 100000
class PermMissingElemTest(unittest.TestCase):
    def test_example(self):
        self.assertEqual(4, solution([2, 3, 1, 5]))

    def test_random(self):
        arr = [n for n in range(1, random.randint(0, INT_RANGE))]        
        missing =  random.randint(0, len(arr))
        arr.remove(missing)
        self.assertEqual(missing, solution(arr))

    def text_max(self):
        arr = [n for n in range(1, INT_RANGE)]
        missing = arr.pop()
        self.assertEqual(missing, solution(arr))
