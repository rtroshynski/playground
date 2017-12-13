// Package fibonacci provides a few different implementations of the
// Fibonacci function
// source: https://sourcegraph.com/github.com/yon/go-fibonacci/-/blob/fibonacci.go
package fibonacci

// Iterative returns the Fibonacci value of `n`
func Iterative(n int) int {
	curr, next := 0, 1

	for i := 0; i < n; i++ {
		temp := curr + next
		curr = next
		next = temp
	}

	return curr
}

// Recursive returns the Fibonacci value of `n`
func Recursive(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Recursive(n - 1) + Recursive(n - 2)
	}
}

// TailRecursive returns the Fibonacci value of `n`
func TailRecursive(n int) int {
	return tail(n, 0, 1)
}

func tail(n, curr, next int) int {
	switch n {
	case 0:
		return curr
	default:
		return tail(n - 1, next, curr + next)
	}
}
