package challenges

// How many ways are there to form a space crew.
// There were many countries participating in the mission and each country has trained
// several astronauts, but could only delegate some of them to join the crew.

// First of all, how many ways are there to choose astronauts from one country? That’s easy,
// the answer is binomial coefficient. The choice that one country makes is independent from
// other countries’ choices, hence, in order to obtain the total number of possible choices, we
// just have to multiply the number of choices per country. In other words, we have to calculate
// the following number:

// B(T[0], D[0]) · B(T[1], D[1]) · · · · · B(T[N − 1], D[N − 1])

// Values of binomial coefficient can be calculated using the following formula involving
// factorial. The factorial of a non-negative integer n is denoted by n! and is defined as
// follows:
// B(N, K) = N!/(K! · (N − K)!)
// n! = 1 · 2 · 3 · · · · · (n − 1) · n

// For example, 5! = 1 · 2 · 3 · 4 · 5 = 120.
//   B(n, k) = n! / (k! (n-k)!)
//   B(6, 1) = 6! / (1! (6-1)!) = 6! / (1! 5!) = 6
//   B(6, 3) = 6! / (3! (6-3)!) = 6! / (3! 3!) = 20
//   B(6, 4) = 6! / (4! (6-4)!) = 6! / (4! 2!) = 15
//   and so forth

func SolutionSpaceCrews(T []int, D []int) int {
	factorials := make([]int, 1000000)

	var factorial func(val int) int
	factorial = func(val int) int {
		if val == 0 {
			return 1
		}

		if factorials[val] == 0 {
			factorials[val] = val * factorial(val-1)
		}
		return factorials[val]
	}

	calcCoefficient := func(astronautsInCountryCount, delegatedByCountryCount int) int {
		// xCy = n! / (k! (n-k)!)
		n := factorial(astronautsInCountryCount)
		k := factorial(delegatedByCountryCount)
		z := factorial(astronautsInCountryCount - delegatedByCountryCount)
		c := n / (k * z)
		return c
	}

	total := 1
	for i := 0; i < len(T); i++ {
		total *= calcCoefficient(T[i], D[i])
	}
	if total > 1410000016 {
		total %= 1410000017
	}
	return total
}
