package main

import (
	"math"
	"strconv"
	"strings"
)

var data []string

type Node struct {
	Value       int
	Right, Left *Node
}

type BST struct {
	level int
	deep  string
	Tree  *Node
}

func LoadNode(data []int) *Node {
	node := Node{}
	for k, v := range data {
		if k == 0 {
			node.Value = v
		}
		node.Insert(v)
	}
	return &node
}
func (n *Node) Insert(value int) {
	if n.Value < value {
		insertRight(n, value)
	} else if n.Value > value {
		insertLeft(n, value)
	}
}

func insertRight(n *Node, value int) {
	if n.Right != nil {
		n.Right.Insert(value)
	} else {
		n.Right = &Node{Value: value}
	}
}

func insertLeft(n *Node, value int) {
	if n.Left != nil {
		n.Left.Insert(value)
	} else {
		n.Left = &Node{Value: value}
	}
}

func (n *Node) SearchByKey(key int) bool {
	if n == nil {
		return false
	}
	if n.Value < key {
		return n.Right.SearchByKey(key)
	} else if n.Value > key {
		return n.Left.SearchByKey(key)
	}
	return true
}

func (bst *BST) findDeep(node *Node, level int) {
	if node == nil {
		return
	}

	if level == 0 {
		data = append(data, strconv.Itoa(node.Value))
		bst.deep = strings.Join(data, ", ")
	}
	bst.findDeep(node.Left, level-1)
	bst.findDeep(node.Right, level-1)
}
func (bst *BST) Find() (deep string, deepest int) {
	data = nil
	bst.level = calculateLevel(bst.Tree) - 1
	bst.findDeep(bst.Tree, bst.level)
	deep = bst.deep
	deepest = bst.level
	return
}

func calculateLevel(n *Node) int {
	if n == nil {
		return 0
	}
	leftLevel := calculateLevel(n.Left) + 1
	rightLeve := calculateLevel(n.Right) + 1
	return int(math.Max(float64(leftLevel), float64(rightLeve)))
}
