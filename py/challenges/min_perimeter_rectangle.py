
import unittest
import math


# An integer N is given, representing the area of some rectangle.

# The area of a rectangle whose sides are of length A and B is A * B, and the perimeter is 2 * (A + B).

# The goal is to find the minimal perimeter of any rectangle whose area equals N. 
# The sides of this rectangle should be only integers.

# For example, given integer N = 30, rectangles of area 30 are:

# (1, 30), with a perimeter of 62,
# (2, 15), with a perimeter of 34,
# (3, 10), with a perimeter of 26,
# (5, 6), with a perimeter of 22.
# Write a function:

# def solution(N)

# that, given an integer N, returns the minimal perimeter of any rectangle whose area is exactly equal to N.

# For example, given an integer N = 30, the function should return 22, as explained above.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [1..1,000,000,000].



def solution(N):

    area = N
    if area == 1:
        return 4
    
    minPerim = 0
    for w in range (1, area):
        if area % w == 0:
            h = area / w
            perim = (w+h)*2
            if perim < minPerim or minPerim is 0:
                minPerim = perim
    
    return minPerim


class TestMinPerimeterRectangle(unittest.TestCase):
    def test_ex(self):
        self.assertEqual(solution(30), 22)

    def test_min(self):
        self.assertEqual(solution(1), 4)
