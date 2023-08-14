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

	insertionSort(s)

	fmt.Printf("%+v\n", s)
}

func insertionSort(s []int) {
	insert := func(idx int) {
		for k := 0; k < idx; k++ {
			if s[k] > s[idx] {
				s[k], s[idx] = s[idx], s[k]
			}
		}
	}

	for idx := 1; idx < len(s); idx++ {
		insert(idx)
	}
}
