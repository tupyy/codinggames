package main

// implementation of BST

import "fmt"

type Node struct {
	Key         string
	Value       interface{}
	Left, Right *Node

	// nodes in subtree rooted her
	N int
}

func size(n *Node) int {
	if n == nil {
		return 0
	}

	return n.N
}

func put(n *Node, key string, value interface{}) *Node {
	if n == nil {
		return &Node{Key: key, Value: value, N: 1}
	}

	if n.Key > key {
		n.Left = put(n.Left, key, value)
	} else if n.Key < key {
		n.Right = put(n.Right, key, value)
	} else {
		n.Value = value
		n.N = size(n.Left) + size(n.Right) + 1
	}

	return n
}

func get(n *Node, k string) interface{} {
	if n == nil {
		return nil
	}

	if n.Key < k {
		return get(n.Left, k)
	} else if n.Key > k {
		return get(n.Right, k)
	}

	return n.Value

}

func minn(n *Node) string {
	if n == nil {
		return ""
	}

	if n.Left == nil {
		return n.Key
	}

	return minn(n.Left)
}

func maxx(n *Node) string {
	if n == nil {
		return ""
	}

	if n.Right == nil {
		return n.Key
	}

	return maxx(n.Right)
}

func rank(n *Node, r int) string {
	if n == nil {
		return ""
	}

	return ""
}

func deleteMin(n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.Left != nil {
		n.Left = deleteMin(n.Left)
	} else {
		return n.Right
	}

	return n
}

func print(n *Node) {
	if n == nil {
		return
	}

	fmt.Printf("Node: %p, Left: %p, Right: %p, Key: %s, Value: %v, Size: %d\n", n, n.Left, n.Right, n.Key, n.Value, n.N)

	print(n.Left)
	print(n.Right)
}

func main() {
	root := put(nil, "s", 2)
	put(root, "e", 1)
	put(root, "a", 2)
	put(root, "r", 3)
	put(root, "c", 4)
	put(root, "h", 5)
	put(root, "e", 5)
	put(root, "x", 5)
	put(root, "m", 9)

	print(root)
	fmt.Printf("Min: %s\n", minn(root))
	fmt.Printf("Max: %s\n", maxx(root))

	deleteMin(root)
	print(root)
}
