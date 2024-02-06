package challenges

// A prime is a positive integer X that has exactly two distinct divisors: 1 and X.
// The first few prime integers are 2, 3, 5, 7, 11 and 13.

// A semiprime is a natural number that is the product of two (not necessarily distinct)
// prime numbers. The first few semiprimes are 4, 6, 9, 10, 14, 15, 21, 22, 25, 26.

// You are given two non-empty arrays P and Q, each consisting of M integers.
// These arrays represent queries about the number of semiprimes within specified ranges.

// Query K requires you to find the number of semiprimes within the range
// (P[K], Q[K]), where 1 ≤ P[K] ≤ Q[K] ≤ N.

// For example, consider an integer N = 26 and arrays P, Q such that:
//     P[0] = 1    Q[0] = 26
//     P[1] = 4    Q[1] = 10
//     P[2] = 16   Q[2] = 20
// The number of semiprimes within each of these ranges is as follows:

// (1, 26) is 10,
// (4, 10) is 4,
// (16, 20) is 0.
// Write a function:

// func Solution(N int, P []int, Q []int) []int

// that, given an integer N and two non-empty arrays P and Q consisting
// of M integers, returns an array consisting of M elements specifying
// the consecutive answers to all the queries.

// For example, given an integer N = 26 and arrays P, Q such that:
//     P[0] = 1    Q[0] = 26
//     P[1] = 4    Q[1] = 10
//     P[2] = 16   Q[2] = 20
// the function should return the values [10, 4, 0], as explained above.

// Write an efficient algorithm for the following assumptions:

// N is an integer within the range [1..50,000];
// M is an integer within the range [1..30,000];
// each element of arrays P and Q is an integer within the range [1..N];
// P[i] ≤ Q[i].

func SieveOfEratosthenes(upperbound int) []int {
	// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
	primesMap := make([]bool, upperbound+1)
	primes := []int{}
	for i := 0; i < upperbound+1; i++ {
		primesMap[i] = true
	}
	p := 2
	for p*p <= upperbound {
		if primesMap[p] {
			for i := p * p; i <= upperbound; i += p {
				primesMap[i] = false
			}
		}
		p += 1
	}

	// Extract all primes
	for startPrime := 2; startPrime <= upperbound; startPrime++ {
		if primesMap[startPrime] {
			primes = append(primes, startPrime)
		}
	}

	return primes
}

func SolutionCountSemiprimes(N int, P []int, Q []int) []int {

	primes := SieveOfEratosthenes(N)

	semiprimes := make(map[int]bool)
	for _, i := range primes {
		for _, j := range primes {
			if i*j <= N {
				semiprimes[i*j] = true
			}
		}
	}

	primeCount := []int{}
	for i := 0; i < len(P); i++ {
		semiprimesInRange := []int{}
		for k, _ := range semiprimes {
			if k >= P[i] && k <= Q[i] {
				semiprimesInRange = append(semiprimesInRange, k)
			}
		}
		primeCount = append(primeCount, len(semiprimesInRange))
	}

	return primeCount
}
