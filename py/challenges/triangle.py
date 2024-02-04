import unittest
import random

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

