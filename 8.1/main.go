package main

func tripleStep(n int) int {
	if n == 0 {
		return 0
	}
	return tripleStepR(n, make([]int, n+1))
}

func tripleStepR(n int, memo []int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if memo[n] == 0 {
		memo[n] = tripleStepR(n-1, memo) + tripleStepR(n-2, memo) + tripleStepR(n-3, memo)
	}
	return memo[n]
}
