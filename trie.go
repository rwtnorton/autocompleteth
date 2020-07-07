package main

import (
	"bytes"
	"fmt"
	"sort"
)

type Trie struct {
	Root  *Node
	Runes map[rune][]*Node
	Words []*Node
}

func NewTrie() *Trie {
	return &Trie{
		Root:  NewNode('\u0000'),
		Runes: make(map[rune][]*Node),
		Words: []*Node{},
	}
}

func (tr *Trie) Insert(s string) *Node {
	runes := []rune(s)
	currNode := tr.Root
	for _, r := range runes {
		var isNew bool
		currNode, isNew = currNode.AddChild(r)
		if isNew {
			if tr.Runes[r] == nil {
				tr.Runes[r] = []*Node{}
			}
			tr.Runes[r] = append(tr.Runes[r], currNode)
		}
	}
	currNode.Count += 1
	tr.Words = append(tr.Words, currNode)
	return currNode
}


func (tr *Trie) showRunes() string {
	var buf bytes.Buffer
	buf.WriteString("{ ")
	var keys []rune
	for r := range tr.Runes {
		keys = append(keys, r)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for i, r := range keys {
		nodes := tr.Runes[r]
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("[%c: ", r))
		for _, n := range nodes {
			buf.WriteString(fmt.Sprintf(" %p", n))
		}
		buf.WriteString("]")
	}
	buf.WriteString(" }")
	return buf.String()
}
