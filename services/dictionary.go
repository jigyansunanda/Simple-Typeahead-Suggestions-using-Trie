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
