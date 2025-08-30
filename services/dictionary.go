package services

import (
	"container/heap"

	"github.com/jigyansunanda/Simple-Typeahead-Suggestions-using-Trie/structures"
)

const MAX_SUGGESTION_COUNT = 5

type Dictionary struct {
	trie          structures.Trie
	priorityQueue structures.MyHeap
}

func NewDictionary() Dictionary {
	dictionary := &Dictionary{
		trie: *structures.NewTrie(),
	}

	return *dictionary
}

func (dictionary *Dictionary) InsertWordIntoDictionary(word string) {
	dictionary.trie.InsertWord(word)
}

func (dictionary *Dictionary) SearchPrefixInDictionary(word string) bool {
	return dictionary.trie.SearchPrefix(word)
}

func (dictionary *Dictionary) getSuggestions(prefix string, myHeap *structures.MyHeap) {
	currentNode := dictionary.trie.Root
	for _, char := range prefix {
		str := string(char)
		currentNode = currentNode.Child[str]
	}
	for str := range currentNode.Child {
		getSuggestionsUtil(currentNode.Child[str], prefix+str, myHeap)
	}
}

func getSuggestionsUtil(node *structures.Node, word string, myHeap *structures.MyHeap) {
	if node == nil {
		return
	}
	if node.IsEndOfWord {
		if myHeap.Len() < MAX_SUGGESTION_COUNT {
			heap.Push(myHeap, &structures.FrequencyPair{Word: word, Frequency: node.Frequency})
		} else {
			if myHeap.Peek().(*structures.FrequencyPair).Frequency < node.Frequency {
				heap.Pop(myHeap)
				heap.Push(myHeap, &structures.FrequencyPair{Word: word, Frequency: node.Frequency})
			} else if myHeap.Peek().(*structures.FrequencyPair).Frequency == node.Frequency {
				if myHeap.Peek().(*structures.FrequencyPair).Word < word {
					heap.Pop(myHeap)
					heap.Push(myHeap, &structures.FrequencyPair{Word: word, Frequency: node.Frequency})
				}
			}
		}
	}
	for str := range node.Child {
		getSuggestionsUtil(node.Child[str], word+str, myHeap)
	}
}
