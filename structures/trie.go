package structures

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	trie := &Trie{Root: NewNode()}

	return trie
}

func (trie *Trie) InsertWord(word string) {
	currentNode := trie.Root
	for _, char := range word {
		// Convert rune to string
		str := string(char)
		// Check if node exists or not
		childNode, isExist := currentNode.Child[str]
		if !isExist {
			currentNode.Child[str] = NewNode()
			childNode = currentNode.Child[str]
		}
		currentNode = childNode
	}
	currentNode.IsEndOfWord = true
	currentNode.Frequency += 1
}

func (trie *Trie) SearchWord(word string) bool {
	currentNode := trie.Root
	for _, char := range word {
		// Convert rune to string
		str := string(char)
		// Check if node exists or not
		childNode, isExist := currentNode.Child[str]
		if !isExist {
			return false
		}
		currentNode = childNode
	}
	return currentNode.IsEndOfWord
}

func (trie *Trie) SearchPrefix(word string) bool {
	currentNode := trie.Root
	for _, char := range word {
		// Convert rune to string
		str := string(char)
		// Check if node exists or not
		childNode, isExist := currentNode.Child[str]
		if !isExist {
			return false
		}
		currentNode = childNode
	}
	if currentNode.IsEndOfWord {
		currentNode.Frequency += 1
	}
	return true
}
