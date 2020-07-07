package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
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

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func (tr *Trie) MostFrequentWordsMatching(term string, count int) []WordCount {
	if count <= 0 {
		return nil
	}
	runes := []rune(term)
	if len(runes) == 0 {
		return nil
	}
	candidates := tr.Runes[runes[0]]
	var matches []*Node
	for _, n := range candidates {
		matchingWordNodes := n.MatchTerm(term)
		if matchingWordNodes == nil {
			continue
		}
		matches = append(matches, matchingWordNodes...)
	}
	return mapWordNodesIntoWordCounts(matches, count)
}

func (tr *Trie) MostFrequentWords(count int) []WordCount {
	if count <= 0 {
		return nil
	}
	return mapWordNodesIntoWordCounts(tr.Words, count)
}

func mapWordNodesIntoWordCounts(rawNodes []*Node, count int) []WordCount {
	if rawNodes == nil {
		return nil
	}
	uniqNodes := map[*Node]struct{}{}
	for _, n := range rawNodes {
		uniqNodes[n] = struct{}{}
	}
	var wordNodes []*Node
	for k := range uniqNodes {
		wordNodes = append(wordNodes, k)
	}
	sort.SliceStable(wordNodes, func(i, j int) bool {
		if wordNodes[i].Count == wordNodes[j].Count {
			return strings.Compare(
				string(wordNodes[i].Word()),
				string(wordNodes[j].Word())) < 0
		}
		return wordNodes[i].Count > wordNodes[j].Count
	})
	var results []WordCount
	for i, node := range wordNodes {
		if i >= count {
			break
		}
		results = append(results, WordCount{string(node.Word()), node.Count})
	}
	return results
}
