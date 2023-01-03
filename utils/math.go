package utils

func CalcPermutation(n, r int) int {
	if n < r {
		return 0
	}
	var sum int = 1
	for i := (n - r + 1); i <= n; i++ {
		sum *= i
		// logger.DebugF("sum = %d * %d\n", sum, i)
	}
	return sum
}

func CalcPermutationMore(n, r1, r2 int) int {
	var total int
	for r := r1; r <= r2; r++ {
		total += CalcPermutation(n, r)
	}
	return total
}
