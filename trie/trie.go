package trie

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

func (trie *Trie) Insert(s string) *Node {
	runes := []rune(s)
	currNode := trie.Root
	for _, r := range runes {
		currNode, isNew := currNode.AddChild(r)
		if isNew {
			nodes, found := trie.Runes[r]
			if !found {
				trie.Runes[r] = []*Node{currNode}
			} else {
				nodes = append(nodes, currNode)
			}
		}
	}
	currNode.Count += 1
	return currNode
}
