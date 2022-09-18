package main

import (
	"fmt"
	avltree "github.com/pristupaanastasia/lexArray/avltree"
	"strings"
)

type AssocArray interface {
	Assign(s string, x int)
	Lookup(s string) (x int, exist bool)
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
	var s string
	tree = &avltree.AvlNode{}
	s = "alpha x1 beta alpha beta ggg"
	fmt.Println(s)
	fmt.Println(lex(s, tree))
	s2 := "    alpha x1     beta     alpha beta     ggg    1  56 "
	fmt.Println(s2)
	fmt.Println(lex(s2, tree))
	s3 := "    a b c d a b d g g k m b "
	fmt.Println(s3)
	fmt.Println(lex(s3, tree))
}
