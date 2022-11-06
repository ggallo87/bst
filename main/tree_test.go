package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	list = []int{12, 11, 90, 82, 7, 9}
	n    *Node
	bst  BST
)

func TestMain(m *testing.M) {
	n = LoadNode(list)
	bst = BST{Tree: n}
	os.Exit(m.Run())
}
func TestLoadNode(t *testing.T) {
	//Node is loaded by main function
	assert.NotNil(t, n)
}

func TestNode_GetLevel(t *testing.T) {
	level := n.GetLevel()
	assert.Equal(t, 3, level)
}

func TestNode_SearchByKey(t *testing.T) {
	existNumber := n.SearchByKey(11)
	assert.True(t, existNumber)
}

func TestNode_SearchByKeyNotExisst(t *testing.T) {
	existNumber := n.SearchByKey(20)
	assert.False(t, existNumber)
}

func TestMinValue(t *testing.T) {
	minNode := MinValue(n)
	assert.Equal(t, 7, minNode.Value)
}

func TestBST_Find(t *testing.T) {
	deep, deepest := bst.Find()
	assert.Equal(t, "9", deep)
	assert.Equal(t, 3, deepest)
}

func TestBST_GetRouteInOrderNode(t *testing.T) {
	order := bst.GetRouteInOrderNode()
	assert.Equal(t, "7,9,11,12,82,90", order)
}

func TestBST_GetRoutePreOrderNode(t *testing.T) {
	order := bst.GetRoutePreOrder()
	assert.Equal(t, "12,11,7,9,90,82", order)
}

func TestBST_GetRoutePostOrderNode(t *testing.T) {
	order := bst.GetRoutePostOrder()
	assert.Equal(t, "11,7,9,90,82,12", order)
}

func TestRemoveValue(t *testing.T) {
	newNode := RemoveValue(bst.Tree, 9)
	assert.NotNil(t, newNode)
	assert.False(t, newNode.SearchByKey(9))

	newNode2 := RemoveValue(newNode, 90)
	assert.NotNil(t, newNode2)
	assert.False(t, newNode2.SearchByKey(20))

}

func TestRemoveValueNotExistValue(t *testing.T) {
	newNode := RemoveValue(bst.Tree, 99)
	assert.Equal(t, bst.Tree, newNode)
}
