
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

