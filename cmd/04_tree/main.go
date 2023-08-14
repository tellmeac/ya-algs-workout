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
		Val: 1,
		Left: &Node{
			Val: 2,
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
				},
				Right: &Node{
					Val: 3,
				},
			},
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
