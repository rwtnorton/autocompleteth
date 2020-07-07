package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Rune     rune
	Children map[rune]*Node
	Parent   *Node
	Count    int
}

func NewNode(r rune) *Node {
	return &Node{
		Rune:     r,
		Children: make(map[rune]*Node),
	}
}

func (node *Node) String() string {
	return fmt.Sprintf("[%c, %d]", node.Rune, node.Count)
}

func (node *Node) OrderedChildren() []*Node {
	var keys []rune
	for r := range node.Children {
		keys = append(keys, r)
	}
	var results = make([]*Node, len(keys), len(keys))
	sort.SliceStable(keys, func(i int, j int) bool {
		return keys[i] < keys[j]
	})
	for i, k := range keys {
		results[i] = node.Children[k]
	}
	return results
}

func (node *Node) AddChild(r rune) (*Node, bool) {
	child, found := node.Children[r]
	if found {
		return child, false
	}
	newNode := NewNode(r)
	newNode.Parent = node
	node.Children[r] = newNode
	return newNode, true
}

func (node *Node) Word() []rune {
	var results []rune
	currNode := node
	for currNode != nil && currNode.Rune != '\u0000' {
		results = append(results, currNode.Rune)
		currNode = currNode.Parent
	}
	// Reverse this slice in-place.
	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}
	return results
}
