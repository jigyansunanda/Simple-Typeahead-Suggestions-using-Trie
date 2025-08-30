package services

import "github.com/jigyansunanda/Simple-Typeahead-Suggestions-using-Trie/structures"

type Dictionary struct {
	trie          structures.Trie
	priorityQueue structures.MyHeap
}

func NewDictionary() Dictionary {
	dictionary := &Dictionary{
		trie:          *structures.NewTrie(),
		priorityQueue: structures.MyHeap{},
	}

	return *dictionary
}

func (dictionary *Dictionary) InsertWordIntoDictionary(word string) {
	dictionary.trie.InsertWord(word)
}

func (dictionary *Dictionary) SearchPrefixInDictionary(word string) bool {
	return dictionary.trie.SearchPrefix(word)
}
