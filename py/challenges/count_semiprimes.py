import unittest
import math

# A prime is a positive integer X that has exactly two distinct divisors: 1 and X. 
# The first few prime integers are 2, 3, 5, 7, 11 and 13.

# A semiprime is a natural number that is the product of two (not necessarily distinct) 
# prime numbers. The first few semiprimes are 4, 6, 9, 10, 14, 15, 21, 22, 25, 26.

# You are given two non-empty arrays P and Q, each consisting of M integers. 
# These arrays represent queries about the number of semiprimes within specified ranges.

# Query K requires you to find the number of semiprimes within the range 
# (P[K], Q[K]), where 1 ≤ P[K] ≤ Q[K] ≤ N.

# For example, consider an integer N = 26 and arrays P, Q such that:
#     P[0] = 1    Q[0] = 26
#     P[1] = 4    Q[1] = 10
#     P[2] = 16   Q[2] = 20
# The number of semiprimes within each of these ranges is as follows:

# (1, 26) is 10,
# (4, 10) is 4,
# (16, 20) is 0.
# Write a function:

# class Solution { public int[] solution(int N, int[] P, int[] Q); }
# that, given an integer N and two non-empty arrays P and Q consisting 
# of M integers, returns an array consisting of M elements specifying 
# the consecutive answers to all the queries.

# For example, given an integer N = 26 and arrays P, Q such that:
#     P[0] = 1    Q[0] = 26
#     P[1] = 4    Q[1] = 10
#     P[2] = 16   Q[2] = 20
# the function should return the values [10, 4, 0], as explained above.

# Write an efficient algorithm for the following assumptions:

# N is an integer within the range [1..50,000];
# M is an integer within the range [1..30,000];
# each element of arrays P and Q is an integer within the range [1..N];
# P[i] ≤ Q[i].

MAX = 50000

def solution(N, P, Q):

    def SieveOfEratosthenes():
        upperbound = N
        # https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
        primes_map = [True] * (upperbound + 1)
        primes = []
        p = 2
        while (p * p <= upperbound):    
            if (primes_map[p] == True):    
                for i in range(p * p, upperbound+1, p):
                    primes_map[i] = False
            p += 1

        # Extract all primes
        for start_prime_number in range(2, upperbound + 1):
            if primes_map[start_prime_number]:
                primes.append(start_prime_number)

        return primes
        
    primes = SieveOfEratosthenes()
    
    semiprimes = set()
    for i in primes:
        for j in primes:
            if i * j <= N:
                semiprimes.add(i * j)
    
    primecount = []
    for i, j in zip(P, Q):
        semiprimes_in_range = [x for x in semiprimes if i <= x <= j]        
        primecount.append(len(semiprimes_in_range))
    
    return primecount

class TestCountSemiprimes(unittest.TestCase):
    def test_ex(self):
        self.assertEqual(solution(26, [1, 4, 16], [26, 10, 20]), [10, 4, 0])

    def test_min(self):
        self.assertEqual(solution(1, [1], [1]), [0])

    def test_max(self):
        self.assertEqual(solution(50000, [1], [50000]), [12110])

