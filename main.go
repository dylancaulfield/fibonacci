package main

func fibRecursive(n int) int {

	if n < 2 {
		return 1
	}

	return fibRecursive(n-1) + fibRecursive(n-2)

}

var memo = map[int]int{}

func fibRecursiveMemo(n int) int {

	if n < 2 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}

	fib := fibRecursiveMemo(n-1) + fibRecursiveMemo(n-2)

	memo[n] = fib

	return fib

}

func fibInOrder(n int) int {

	var fibs []int

	fibs = append(fibs, 1)

	previous := 1
	current := 1

	for i := 1; true; i++ {

		if len(fibs) == n {
			return fibs[n-1]
		}

		x := current

		current = previous + current

		previous = x

		fibs = append(fibs, current)
	}

	return 0

}
