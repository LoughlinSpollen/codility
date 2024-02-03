
# A binary gap within a positive integer N is any maximal sequence of consecutive 
# zeros that is surrounded by ones at both ends in the binary representation of N.

# For example, number 9 has binary representation 1001 and contains a binary gap 
# of length 2. The number 529 has binary representation 1000010001 and contains 
# two binary gaps: one of length 4 and one of length 3. The number 20 has binary 
# representation 10100 and contains one binary gap of length 1. The number 15 has
#  binary representation 1111 and has no binary gaps. The number 32 has binary 
# representation 100000 and has no binary gaps.

# Write a function:

# def solution(N)

# that, given a positive integer N, returns the length of its longest binary gap. 
# The function should return 0 if N doesn't contain a binary gap.

# For example, given N = 1041 the function should return 5, because N has binary representation 
# 10000010001 and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [1..2,147,483,647].
# Copyright 2009–2024 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.


import unittest


def solution(N):

    binaryNum = format(N, 'b')
    longest = 0
    currentGapLen = 0
    inGap = False

    max = len(binaryNum)
    for i in range (0, max):
        if inGap:
            if binaryNum[i] == '0':
                currentGapLen += 1
            else: # gap ends
                if longest < currentGapLen:
                    longest = currentGapLen
                currentGapLen = 0

        if binaryNum[i] == '1' and i+1 < max and binaryNum[i+1] == '0':
            inGap = True
    
    return longest


class BinaryGapTest(unittest.TestCase):

    def test_ex(self):
        res = solution(529)
        self.assertEqual(res, 4)

        res = solution(32)
        self.assertEqual(res, 0)    

        res = solution(1041)
        self.assertEqual(res, 5)  