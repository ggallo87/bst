package main

import "fmt"

func main() {
	list := []int{12, 11, 90, 82, 7, 9}

	n := LoadNode(list)

	tree := BST{Tree: n}
	tree.Find()
	fmt.Println("deep: " + tree.deep)
	fmt.Println(fmt.Sprintf("deepest: %v", tree.level))

	list1 := []int{26, 82, 16, 92, 33}
	n1 := LoadNode(list1)
	tree1 := BST{Tree: n1}
	tree1.Find()
	fmt.Println("deep: " + tree1.deep)
	fmt.Println(fmt.Sprintf("deepest: %v", tree1.level))
}
