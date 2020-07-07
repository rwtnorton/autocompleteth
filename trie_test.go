package main

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

	{
		trie.Insert("foo")

		fNode, found := trie.Root.Children['f']
		if !found || fNode == nil {
			t.Errorf("expected top-level 'f' Node, found=%t, fNode=%v", found, fNode)
		}
		oNode1, found := fNode.Children['o']
		if !found || oNode1 == nil {
			t.Errorf("expected child 'o' Node, found=%t, oNode=%v", found, oNode1)
		}
		oNode2, found := oNode1.Children['o']
		if !found || oNode2 == nil {
			t.Errorf("expected grandchild 'o' Node, found=%t, oNode=%v", found, oNode2)
		}
	}

}

func TestRunes(t *testing.T) {
	tr := NewTrie()
	tr.Insert("foo")
	tr.Insert("bar")
	tr.Insert("baz")

	// fmt.Println(tr.showRunes())

	testCases := []struct{
		r rune
		count int
	}{
		{'a', 1},
		{'b', 1},
		{'f', 1},
		{'o', 2},
		{'r', 1},
		{'z', 1},
	}
	for i, testCase := range testCases {
		if len(tr.Runes[testCase.r]) != testCase.count {
			t.Errorf("%d: expected Runes[%c] be have len %d, got %d", i, testCase.r, testCase.count, len(tr.Runes[testCase.r]))
		}
		for _, runeNode := range tr.Runes[testCase.r] {
			if runeNode.Rune != testCase.r {
				t.Errorf("%d: expected Node to have Rune %c, got %c", i, testCase.r, runeNode.Rune)
			}
		}
	}
}

func TestWords(t *testing.T) {
	tr := NewTrie()
	nodes := []*Node{
		tr.Insert("foo"),
		tr.Insert("bar"),
		tr.Insert("baz"),
	}

	for i, node := range nodes {
		if tr.Words[i] != node {
			t.Errorf("%d: expected Word node %p, got %p", i, node, tr.Words[i])
		}
		if tr.Words[i].Count <= 0 {
			t.Errorf("%d: expected Count of Word node to be > 0, got %d", i, tr.Words[i].Count)
		}
	}
}
