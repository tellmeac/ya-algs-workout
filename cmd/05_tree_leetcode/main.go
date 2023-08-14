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

	required := 22

	ok := false
	dfs(root, 0, required, func() {
		ok = true
	})

	fmt.Println(ok)
}

// dfs, use validate to
func dfs(node *Node, acc, req int, validate func()) {
	if node == nil {
		return
	}

	acc = acc + node.Val

	if node.Left == nil && node.Right == nil && acc == req {
		validate()
	}

	dfs(node.Left, acc, req, validate)
	dfs(node.Right, acc, req, validate)
}
