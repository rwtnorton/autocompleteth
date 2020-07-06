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

func (tr *Trie) Insert(s string) *Node {
	runes := []rune(s)
	currNode := tr.Root
	for _, r := range runes {
		var isNew bool
		currNode, isNew = currNode.AddChild(r)
		if isNew {
			nodes, found := tr.Runes[r]
			if !found {
				tr.Runes[r] = []*Node{currNode}
			} else {
				nodes = append(nodes, currNode)
			}
		}
	}
	currNode.Count += 1
	tr.Words = append(tr.Words, currNode)
	return currNode
}
