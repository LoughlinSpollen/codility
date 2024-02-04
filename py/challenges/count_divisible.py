import unittest

# Write a function:

# def solution(A, B, K)

# that, given three integers A, B and K, returns the number of integers 
# within the range [A..B] that are divisible by K, i.e.:

# { i : A ≤ i ≤ B, i mod K = 0 }

# For example, for A = 6, B = 11 and K = 2, your function should return 3,
# because there are three numbers divisible by 2 within the range [6..11], namely 6, 8 and 10.

# Write an efficient algorithm for the following assumptions:

# A and B are integers within the range [0..2,000,000,000];
# K is an integer within the range [1..2,000,000,000];
# A ≤ B.

def solution(A, B, K):
    start = A
    end = B
    divider = K

    result = int(end/divider) - int(start/divider)
    if start % divider == 0:
        return result + 1
    return result


class CountDivTest(unittest.TestCase):
    def test_example(self):
        self.assertEqual(solution(6, 11, 2), 3)

    def test_min(self):
        self.assertEqual(solution(0, 0, 11), 1)
        self.assertEqual(solution(1, 1, 11), 0)

    def test_max(self):
        self.assertEqual(solution(0, 2000000000, 1), 2000000001)
        self.assertEqual(solution(0, 2000000000, 2000000000), 2)
        self.assertEqual(solution(0, 2000000000, 1999999999), 2)

