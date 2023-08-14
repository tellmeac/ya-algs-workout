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

func dfs(node *Node, add func(int)) {
	if node == nil {
		return
	}

	add(node.Val)

	dfs(node.Left, add)
	dfs(node.Right, add)
}
