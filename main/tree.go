package main

import (
	"math"
	"strconv"
	"strings"
)

var (
	data  []string
	route []string
)

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
	if n.Value == value {
		return
	}
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

func (n *Node) GetLevel() int {
	return calculateLevel(n) - 1
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
	bst.level = bst.Tree.GetLevel()
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

func RemoveValue(n *Node, value int) *Node {
	if n == nil {
		return n
	}
	if value < n.Value {
		n.Left = RemoveValue(n.Left, value)
	}
	if value > n.Value {
		n.Right = RemoveValue(n.Right, value)
	}

	if n.Value == value {
		if n.Left == nil && n.Right == nil {
			n = nil
			return n
		}
		if n.Left == nil && n.Right != nil {
			temp := n.Right
			n = nil
			n = temp
			return n
		}
		if n.Left != nil && n.Right == nil {
			temp := n.Left
			n = nil
			n = temp
			return n
		}
		tempNode := MinValue(n.Right)
		n.Value = tempNode.Value
		n.Right = RemoveValue(n.Right, tempNode.Value)
	}
	return n
}

func MinValue(node *Node) *Node {
	temp := node
	for nil != temp && temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

func (bst *BST) GetRouteInOrderNode() string {
	route = nil
	result := new(string)
	inOrder(bst.Tree, result)
	return *result
}

func inOrder(node *Node, res *string) {
	if node != nil {
		inOrder(node.Left, res)
		route = append(route, strconv.Itoa(node.Value))
		inOrder(node.Right, res)
	}
	*res = strings.Join(route, ",")
}

func (bst *BST) GetRoutePreOrder() string {
	route = nil
	result := new(string)
	preOrder(bst.Tree, result)
	return *result
}

func preOrder(node *Node, res *string) {
	if node != nil {
		route = append(route, strconv.Itoa(node.Value))
		preOrder(node.Left, res)
		preOrder(node.Right, res)
	}
	*res = strings.Join(route, ",")
}
func (bst *BST) GetRoutePostOrder() string {
	route = nil
	result := new(string)
	postOrder(bst.Tree, result)
	return *result
}

func postOrder(node *Node, res *string) {
	if node != nil {
		preOrder(node.Left, res)
		preOrder(node.Right, res)
		route = append(route, strconv.Itoa(node.Value))
	}
	*res = strings.Join(route, ",")
}
