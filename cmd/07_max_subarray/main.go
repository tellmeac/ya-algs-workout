package main

import "fmt"

func main() {
	s := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	ans := 0
	current := 0
	for idx := 0; idx < len(s); idx++ {
		current += s[idx]
		ans = max(ans, current)

		if current < 0 {
			current = 0
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
