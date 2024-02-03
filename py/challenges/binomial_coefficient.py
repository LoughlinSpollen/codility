import math
import unittest

# How many ways are there to form a space crew. 
# There were many countries participating in the mission and each country has trained 
# several astronauts, but could only delegate some of them to join the crew.
# 
# First of all, how many ways are there to choose astronauts from one country? That’s easy,
# the answer is binomial coefficient. The choice that one country makes is independent from
# other countries’ choices, hence, in order to obtain the total number of possible choices, we
# just have to multiply the number of choices per country. In other words, we have to calculate
# the following number:
# 
# B(T[0], D[0]) · B(T[1], D[1]) · · · · · B(T[N − 1], D[N − 1])
# 
# Values of binomial coefficient can be calculated using the following formula involving
# factorial. The factorial of a non-negative integer n is denoted by n! and is defined as
# follows:
# B(N, K) = N!/(K! · (N − K)!)
# n! = 1 · 2 · 3 · · · · · (n − 1) · n
#
# For example, 5! = 1 · 2 · 3 · 4 · 5 = 120.
#   B(n, k) = n! / (k! (n-k)!)
#   B(6, 1) = 6! / (1! (6-1)!) = 6! / (1! 5!) = 6
#   B(6, 3) = 6! / (3! (6-3)!) = 6! / (3! 3!) = 20
#   B(6, 4) = 6! / (4! (6-4)!) = 6! / (4! 2!) = 15
#   and so forth

def solution(T, D):
        factorials = [0]*1000000
        
        def calcCoefficient(astronautsInCountryCount, delegatedByCountryCount):
            # xCy = n! / (k! (n-k)!)
            n = factorial(astronautsInCountryCount)
            k = factorial(delegatedByCountryCount)
            z = factorial(astronautsInCountryCount-delegatedByCountryCount)
            c = n // (k * z)  
            return c

        def factorial(val):
            if val == 0:
                return 1

            if factorials[val] == 0:
                factorials[val] = val * factorial(val-1)
            return factorials[val]
        
        astronautsInCountries = T
        delegatedByCountries = D

        total = 1
        for i in range(0, len(astronautsInCountries)):
            total *= calcCoefficient(astronautsInCountries[i], delegatedByCountries[i])
            print(total)
        if total > 1410000016:
            total %= 1410000017
        return total

    
class SpaceCrewTest(unittest.TestCase):

    def test_ex(self):
        T = [6,4,7]
        D = [1,3,4]    
        res = solution(T, D)
        self.assertEqual(res, 840)

    