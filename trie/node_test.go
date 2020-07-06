package trie

import (
	"testing"
)

func TestOrderedChildren(t *testing.T) {
	node := NewNode('n')
	aNode := NewNode('a')
	xNode := NewNode('x')
	mNode := NewNode('m')
	for _, n := range []*Node{aNode, xNode, mNode} {
		node.Children[n.Rune] = n
		n.Parent = node
	}
	got := node.OrderedChildren()
	if len(got) != 3 {
		t.Errorf("expected 3 Nodes, got %d: %v", len(got), got)
	}
	if got[0] != aNode {
		t.Errorf("expected Node with rune 'a' at index 0, got %v", got[0])
	}
	if got[1] != mNode {
		t.Errorf("expected Node with rune 'm' at index 1, got %v", got[1])
	}
	if got[2] != xNode {
		t.Errorf("expected Node with rune 'x' at index 2, got %v", got[2])
	}
}

func TestAddChild(t *testing.T) {
	node := NewNode('a')
	got, isNew := node.AddChild('b')
	if !isNew {
		t.Errorf("expected AddChild on new rune to show isNew true, got %v", isNew)
	}
	if got == nil {
		t.Errorf("expected AddChild to return new *Node, got %v", got)
	}
	if got.Parent != node {
		t.Errorf("expected AddChild to point Parent at caller, got %v", got.Parent)
	}
	if got.Rune != 'b' {
		t.Errorf("expected AddChild to return *Node with rune 'b', got %v", got.Rune)
	}

	got2, isNew2 := node.AddChild('b')
	if isNew2 {
		t.Errorf("expected AddChild on old rune to show isNew false, got %v", isNew2)
	}
	if got2 == nil {
		t.Errorf("expected AddChild to return non-nil *Node, got %v", got2)
	}
	if got2.Parent != node {
		t.Errorf("expected AddChild to point Parent at caller, got %v", got2.Parent)
	}
	if got2.Rune != 'b' {
		t.Errorf("expected AddChild to return *Node with rune 'b', got %v", got2.Rune)
	}
}

func TestWord(t *testing.T) {
	word := "quux"
	runes := []rune(word)
	var nodes []*Node
	for i, r := range runes {
		newNode := NewNode(r)
		nodes = append(nodes, newNode)
		if i == 0 {
			continue
		}
		parent := nodes[i-1]
		newNode.Parent = parent
		parent.Children[r] = newNode
	}

	got := string(nodes[len(nodes)-1].Word())
	if word != got {
		t.Errorf("expected %q, got %q", word, got)
	}
}
