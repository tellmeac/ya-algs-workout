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

	sortSelection(s)

	fmt.Printf("%+v\n", s)
}

func sortSelection(s []int) {
	for idx := 0; idx < len(s); idx++ {
		lowIdx := idx
		for jdx := idx; jdx < len(s); jdx++ {
			if s[lowIdx] > s[jdx] {
				lowIdx = jdx
			}
		}

		s[lowIdx], s[idx] = s[idx], s[lowIdx]
	}
}
