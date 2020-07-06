package trie

import (
	"testing"
)

func TestInsert(t *testing.T) {
	trie := NewTrie()
	{
		trie.Insert("ab")

		aNode, found := trie.Root.Children['a']
		if !found || aNode == nil {
			t.Errorf("expected top-level 'a' Node, found=%t, aNode=%v", found, aNode)
		}
		bNode, found := aNode.Children['b']
		if !found || bNode == nil {
			t.Errorf("expected child 'b' Node, found=%t, bNode=%v", found, bNode)
		}
	}

	{
		trie.Insert("ac")

		aNode, found := trie.Root.Children['a']
		if !found || aNode == nil {
			t.Errorf("expected top-level 'a' Node, found=%t, aNode=%v", found, aNode)
		}
		cNode, found := aNode.Children['c']
		if !found || cNode == nil {
			t.Errorf("expected child 'c' Node, found=%t, bNode=%v", found, cNode)
		}
	}
}
