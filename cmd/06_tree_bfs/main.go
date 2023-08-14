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

	queue := make([]*Node, 0, 100)

	traverse := bfs(root, queue)

	fmt.Println(traverse)
}

func bfs(root *Node, queue []*Node) []int {
	queue = append(queue, root)

	var traverse []int
	var node *Node

	for len(queue) > 0 {
		node, queue = queue[0], queue[1:] // pop from queue

		traverse = append(traverse, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return traverse
}
