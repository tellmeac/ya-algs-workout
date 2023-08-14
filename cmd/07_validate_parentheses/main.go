package main

import "fmt"

func main() {
	fmt.Println(isValid("(){}[]{)["))
}

func isValid(s string) bool {
	// stack contains only opens
	stack := make([]byte, 0, 100)

	data := []byte(s)

	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var lastOpen byte
	for idx := range data {
		_, isOpen := pairs[data[idx]]
		if isOpen {
			stack = append(stack, data[idx])
			continue
		}

		// pop from stack and check if the same
		if len(stack) == 0 {
			return false
		}

		lastOpen, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if pairs[lastOpen] != data[idx] {
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
