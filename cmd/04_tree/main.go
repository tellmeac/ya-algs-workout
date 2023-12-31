package main

import (
	"fmt"
)

type Node struct {
	Val         int
	Left, Right *Node
}

func main() {
	root := &Node{
		Val: 5,
		Left: &Node{
			Val: 4,
			Left: &Node{
				Val: 11,
				Left: &Node{
					Val: 7,
				},
				Right: &Node{
					Val: 2,
				},
			},
			Right: nil,
		},
		Right: &Node{
			Val:   8,
			Left:  nil,
			Right: nil,
		},
	}

	sum := 0

	dfs(root, func(v int) {
		sum += v
	})

	fmt.Println(sum)
}

func dfs(node *Node, callback func(val int)) {
	if node == nil {
		return
	}

	callback(node.Val)

	dfs(node.Left, callback)
	dfs(node.Right, callback)
}
