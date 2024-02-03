import unittest

# A small frog wants to get to the other side of the road. 
# The frog is currently located at position X and wants to get to a position 
# greater than or equal to Y. The small frog always jumps a fixed distance, D.

# Count the minimal number of jumps that the small frog must perform to reach its target.

# Write a function:

# def solution(X, Y, D)

# that, given three integers X, Y and D, returns the minimal number of jumps from 
# position X to a position equal to or greater than Y.


# For example, given:

#   X = 10
#   Y = 85
#   D = 30
# the function should return 3, because the frog will be positioned as follows:

# after the first jump, at position 10 + 30 = 40
# after the second jump, at position 10 + 30 + 30 = 70
# after the third jump, at position 10 + 30 + 30 + 30 = 100
# Write an efficient algorithm for the following assumptions:

# X, Y and D are integers within the range [1..1,000,000,000];
# X ≤ Y.


def solution(X, Y, D):
    startPos = X
    endPos = Y
    jumpLen = D

    j = (endPos - startPos) // jumpLen
    if (endPos - startPos) % jumpLen > 0:
        j += 1
    return j



class FrogJumpTest(unittest.TestCase):
    def test_ex(self):
        self.assertEqual(solution(10, 85, 30), 3)

    def test_max(self):
        MAX_RANGE=1000000000
        self.assertEqual(solution(0, MAX_RANGE, 1), MAX_RANGE)
        self.assertEqual(solution(0, 10, 20), 1)
        self.assertEqual(solution(10, 100, 10), 9)
        self.assertEqual(solution(10, 10, 10), 0)
        self.assertEqual(solution(9, 29, 10), 2)

    def test_min(self):
        self.assertEqual(solution(1, 1, 1), 0)
        self.assertEqual(solution(1, 2, 1), 1)
        self.assertEqual(solution(1, 2, 2), 1)
        self.assertEqual(solution(1, 4, 1), 3)
        self.assertEqual(solution(1, 6, 1), 5)
        