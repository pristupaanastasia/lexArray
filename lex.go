package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AssocArray interface {
	Assign(s string, x int)
	Lookup(s string) (x int, exist bool)
}
type Element struct {
	s string
	x int
}

type AvlNode struct {
	Elem   Element
	Left   *AvlNode
	Right  *AvlNode
	Back   *AvlNode
	Height int
}

func (tree *AvlNode) Assign(s string, x int) {
	if tree.Back != nil {
		tree.Back.Assign(s, x)
	} else {
		tree = tree.Insert(s, x)
	}
}

func (tree *AvlNode) Insert(s string, x int) *AvlNode {
	if tree.Elem.s == "" {
		tree.Elem = Element{s: s, x: x}
		tree.Height = 0
	} else if s < tree.Elem.s {
		if tree.Left == nil {
			tree.Left = &AvlNode{}
			tree.Left.Back = tree
		}
		tree.Left = tree.Left.Insert(s, x)
	} else if s > tree.Elem.s {
		if tree.Right == nil {
			tree.Right = &AvlNode{}
			tree.Right.Back = tree
		}
		tree.Right = tree.Right.Insert(s, x)
	}
	tree = tree.rebalancer()
	tree.Height = tree.findMaxHeight()
	return tree
}

func (tree *AvlNode) Lookup(s string) (x int, exist bool) {
	if tree.Back != nil {
		return tree.Back.Lookup(s)
	} else {
		x, exist = tree.Find(s)
		return x, exist
	}

}

func (tree *AvlNode) Find(s string) (x int, exist bool) {

	if tree.Elem.s == s {
		return tree.Elem.x, true
	} else if s < tree.Elem.s && tree.Left != nil {
		return tree.Left.Find(s)
	} else if s > tree.Elem.s && tree.Right != nil {
		return tree.Right.Find(s)
	} else {
		return 0, false
	}
}

func NodeHeight(tree *AvlNode) int {
	if tree == nil {
		return -1
	} else {
		return tree.Height
	}
}

func (tree *AvlNode) findMaxHeight() int {
	if NodeHeight(tree.Left) > NodeHeight(tree.Right) {
		return NodeHeight(tree.Left) + 1
	} else {
		return NodeHeight(tree.Right) + 1
	}
}

func (tree *AvlNode) rebalancer() *AvlNode {
	if tree == nil {
		return tree
	}
	tree.Height = tree.findMaxHeight()

	balanceFactor := NodeHeight(tree.Left) - NodeHeight(tree.Right)
	if balanceFactor == -2 {
		if NodeHeight(tree.Right.Left) > NodeHeight(tree.Right.Right) {
			tree.Right = LeftRotate(tree.Right)
		}
		return RightRotate(tree)
	} else if balanceFactor == 2 {

		if NodeHeight(tree.Left.Right) > NodeHeight(tree.Left.Left) {
			tree.Left = RightRotate(tree.Left)
		}
		return LeftRotate(tree)
	}
	return tree
}

func LeftRotate(tree *AvlNode) *AvlNode {

	var left *AvlNode
	if tree != nil {
		left = tree.Left
		tree.Left = left.Right
		left.Right = tree

		left.Back = tree.Back
		tree.Back = left
		if tree.Left != nil {
			tree.Left.Back = tree
		}
		//update height
		tree.Height = tree.findMaxHeight()
		left.Height = left.findMaxHeight()
		tree = left
	}
	return tree
}

func RightRotate(tree *AvlNode) *AvlNode {
	var right *AvlNode
	if tree != nil {
		right = tree.Right
		tree.Right = right.Left
		right.Left = tree

		right.Back = tree.Back
		tree.Back = right
		if tree.Right != nil {
			tree.Right.Back = tree
		}

		//update height
		tree.Height = tree.findMaxHeight()
		right.Height = right.findMaxHeight()
		tree = right
	}

	return tree
}
func tookArray(m []string) []string {
	var a []string

	for _, v := range m {
		if v != " " && v != "" && ((v[0] > 64 && v[0] < 91) || (v[0] > 96 && v[0] < 123)) {
			a = append(a, v)
		}
	}
	return a
}

func lex(sentence string, Array AssocArray) []int {
	solve := make([]int, 0)
	array := strings.Split(sentence, " ")
	n := 1
	array = tookArray(array)
	for _, v := range array {
		if x, exist := Array.Lookup(v); exist {
			solve = append(solve, x)
		} else {
			Array.Assign(v, n)
			x, exist = Array.Lookup(v)
			solve = append(solve, x)
			n++
		}
	}
	return solve
}

func main() {

	var tree AssocArray
	var line string
	tree = &AvlNode{}
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line = scanner.Text()
	}

	fmt.Println(lex(line, tree))
}
