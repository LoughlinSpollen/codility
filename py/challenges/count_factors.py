
import unittest
import math

# A positive integer D is a factor of a positive integer N if there exists an integer M such that N = D * M.
# For example, 6 is a factor of 24, because M = 4 satisfies the above condition (24 = 6 * 4).

# Write a function:

# def solution(N)
# that, given a positive integer N, returns the number of its factors.
# For example, given N = 24, the function should return 8, because 24 has 8 factors, 
# namely 1, 2, 3, 4, 6, 8, 12, 24. There are no other factors of 24.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [1..2,147,483,647].

def solution(N):
    num = N
    count = 0      
    sq_root = int(math.sqrt(num))
    for i in range(1, sq_root+1):
        if num % i == 0:
            count += 2
    
    if num % math.sqrt(num) == 0: 
        count -=1

    return count

class TestCountFactors(unittest.TestCase):
    def test_ex(self):
        self.assertEqual(solution(24), 8)
    
    def test_min(self):
        self.assertEqual(solution(1), 1)

    def test_max(self):
        self.assertEqual(solution(2147483647), 2)
        self.assertEqual(solution(2147395600), 135)
        self.assertEqual(solution(2147483646), 4)
        self.assertEqual(solution(2147483645), 8)
        self.assertEqual(solution(2147483644), 12)
        self.assertEqual(solution(2147483643), 4)
        self.assertEqual(solution(2147483642), 4)
    
    def test_extreme(self):
        self.assertEqual(solution(16), 5)
        self.assertEqual(solution(2147483648), 1)
        self.assertEqual(solution(26), 2)
        self.assertEqual(solution(25), 3)
        self.assertEqual(solution(2147483649), 2)
        self.assertEqual(solution(2147483650), 4)
        self.assertEqual(solution(2147483651), 4)
        self.assertEqual(solution(2147483652), 12)
        self.assertEqual(solution(2147483653), 4)
        self.assertEqual(solution(2147483654), 4)
        self.assertEqual(solution(2147483655), 4)
        self.assertEqual(solution(2147483656), 12)
        self.assertEqual(solution(2147483657), 2)
        self.assertEqual(solution(2147483658), 4)
        self.assertEqual(solution(2147483659), 4)
        self.assertEqual(solution(2147483660), 12)
        self.assertEqual(solution(2147483661), 4)
        self.assertEqual(solution(2147483662), 4)
        self.assertEqual(solution(2147483663), 4)
        self.assertEqual(solution(2147483664), 12)
        self.assertEqual(solution(2147483665), 4)
        self.assertEqual(solution(2147483666), 4)
        self.assertEqual(solution(2147483667), 4)
        self.assertEqual(solution(2147483668), 12)
        self.assertEqual(solution(2147483669), 4)
        self.assertEqual(solution(2147483670), 4)
        self.assertEqual(solution(2147483671), 4)
        self.assertEqual(solution(2147483672), 12)
        self.assertEqual(solution(2147483673), 2)
        self.assertEqual(solution(2147483674), 4)
        self.assertEqual(solution(2147483675), 4)
        self.assertEqual(solution(2147483676), 12)
        self.assertEqual(solution(2147483677), 4)
        self.assertEqual(solution(2147483678), 4)
        self.assertEqual(solution(2147483679), 4)
        self.assertEqual(solution(2147483680), 12)
        self.assertEqual(solution(2147483681), 4)
        self.assertEqual(solution(2147483682), 4)
        self.assertEqual(solution(2147483683), 4)


