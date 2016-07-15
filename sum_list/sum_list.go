package sum_list

func SumSeriesRecursive(n int) int {
	if n == 1 {
		return n
	}
	return n + SumSeriesRecursive(n-1)
}

func SumSeriesLoop(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

func SumSeriesFormula(n int) int {
	return (n * (n + 1)) / 2
}
