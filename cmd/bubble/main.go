package main

import (
	"fmt"
)

//var (
//	reader = bufio.NewReader(os.Stdin)
//	writer = bufio.NewWriter(os.Stdout)
//)

func main() {
	s := []int{5, 3, 6, 1}

	sortBubble(s)

	fmt.Printf("%+v", s)
}

func sortBubble(s []int) {
	for idx := 0; idx < len(s)-1; idx++ {
		for jdx := idx + 1; jdx < len(s); jdx++ {
			if s[idx] > s[jdx] {
				s[idx], s[jdx] = s[jdx], s[idx]
			}
		}
	}
}
