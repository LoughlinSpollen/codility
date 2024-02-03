import unittest

# An array A consisting of N integers is given. 
# The dominator of array A is the value that occurs in more than half of the elements of A.

# For example, consider array A such that

#  A[0] = 3    A[1] = 4    A[2] =  3
#  A[3] = 2    A[4] = 3    A[5] = -1
#  A[6] = 3    A[7] = 3
# The dominator of A is 3 because it occurs in 5 out of 8 elements of A 
# (namely in those with indices 0, 2, 4, 6 and 7) and 5 is more than a half of 8.

# Write a function

# def solution(A)

# that, given an array A consisting of N integers, returns index of any element of array A 
# in which the dominator of A occurs. The function should return −1 if array A does not 
# have a dominator.

# For example, given array A such that

#  A[0] = 3    A[1] = 4    A[2] =  3
#  A[3] = 2    A[4] = 3    A[5] = -1
#  A[6] = 3    A[7] = 3
# the function may return 0, 2, 4, 6 or 7, as explained above.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [0..100,000];
# each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].


def solution(A):
    arr = A
    map = dict()

    most_common_k = None
    most_common_v = 0
    for v in arr:
        if v in map:
            map[v] += 1
        else:
            map[v] = 1

        if most_common_k is None or map[v] > most_common_v:
            most_common_k = v
            most_common_v = map[v]
    
    return most_common_k


class TestDominator(unittest.TestCase):
    def test_ex(self):
        self.assertEqual(solution([3,4,3,2,3,-1,3,3]), 3)

    def test_min(self):
        self.assertEqual(solution([1]), 1)

    def test_max(self):
        self.assertEqual(solution([1, 2, 3, 4, 5, 6, 7, 8, 9]), 1)
        self.assertEqual(solution([1, 1, 1, 1, 1, 1, 1, 1, 1]), 1)
        self.assertEqual(solution([1, 1, 1, 1, 1, 1, 1, 1, 2]), 1)

    def test_extreme(self):
        self.assertEqual(solution([1, 2, 3, 4, 5, 6, 7, 8, 9] * 100000), 1)
        
    
    