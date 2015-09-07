package challenge2

import "math"
import "math/big"

// Run calculates the if n is nearly prime
func Run(n int) bool {
	largestPossibleMultiple := int(math.Sqrt(float64(n)))

	for i := 2; i <= largestPossibleMultiple; i++ {
		if n%i != 0 {
			continue
		}

		if !prime(i) {
			continue
		}

		if prime(n / i) {
			return true
		}
	}
	return false
}

func prime(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
}
