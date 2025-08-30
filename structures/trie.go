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
			// Create new node if does not exist
			currentNode.Child[str] = NewNode()
			// Assign the new created node to childnode
			childNode = currentNode.Child[str]
		}
		currentNode = childNode
	}
	// Mark end of word
	currentNode.IsEndOfWord = true
	// Increase word frequency
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
	// To make sure it's not just a prefix
	// but truly an end of the word
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
	// If complete word has been searched, increase frequency
	if currentNode.IsEndOfWord {
		currentNode.Frequency += 1
	}
	return true
}
