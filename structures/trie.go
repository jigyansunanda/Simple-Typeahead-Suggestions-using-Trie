package structures

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	trie := &Trie{Root: NewNode()}

	return trie
}
