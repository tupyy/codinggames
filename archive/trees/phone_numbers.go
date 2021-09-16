package main

import (
	"fmt"
	"strconv"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
func digits(nb string) []int {
	d := []int{}

	for _, s := range nb {
		digit, _ := strconv.Atoi(string(s))
		d = append(d, digit)
	}

	return d
}

type Node struct {
	Parent   *Node
	Children []*Node
	Value    int
}

func findChildrenByValue(value int, root *Node) *Node {
	for _, c := range root.Children {
		if c.Value == value {
			return c
		}
	}

	return nil
}

func count(root *Node) int {
	v := len(root.Children)

	if root.Parent == nil {
		v += 1
	}

	fmt.Printf("%p %+v\n", root, root)
	for _, c := range root.Children {
		v = v + count(c)
	}

	return v
}

func addPhoneNumber(root *Node, digits []int) {
	var n *Node
	for i := 0; i < len(digits); i++ {
		n = &Node{Parent: root, Value: digits[i]}
		root.Children = append(root.Children, n)
		root = n
	}
}

func findInTree(start *Node, value int) *Node {
	s := *start
	ptrS := &s

	found := false
	for {
		if child := findChildrenByValue(value, ptrS); child != nil {
			ptrS = child
			found = true
		} else {
			if !found {
				return nil
			}
			return ptrS
		}
	}
}

func getOrCreate(d int, roots map[int]*Node) *Node {
	if n, found := roots[d]; found {
		return n
	}

	n := &Node{Value: d}
	roots[d] = n

	return n
}

func main() {
	N := 2

	phones := []string{
		"0000123",
		"01",
	}

	roots := make(map[int]*Node)

	for i := 0; i < N; i++ {
		telephone := phones[i]

		digits := digits(telephone)
		startingNode := getOrCreate(digits[0], roots)

		lastIdx := 1
		included := true
		for idx := 1; idx < len(digits); idx++ {
			d := digits[idx]

			n := findChildrenByValue(d, startingNode)
			lastIdx = idx
			if n == nil {
				included = false
				break
			}
			startingNode = n
		}

		if !included {
			addPhoneNumber(startingNode, digits[lastIdx:])
		}
	}

	total := 0
	for _, r := range roots {
		total += count(r)
	}
	fmt.Println(total)
}
