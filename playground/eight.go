package playground

import "fmt"

const maxTrieSize = 26

type TrieNode struct {
	children [maxTrieSize]*TrieNode
}

type RootOfTrie struct {
	root *TrieNode
}

func (r *RootOfTrie) insertIntoTrie(word string) {
	currentNode := r.root
	for index := range word {
		charIndex := word[index] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &TrieNode{}
		} else {
			fmt.Println("character found", string(word[index]))
			currentNode = currentNode.children[charIndex]
		}
	}
}

func (r *RootOfTrie) initTrie() {
	r.root = &TrieNode{}
	// fmt.Println(r.root)
}

func StartTrie() {
	r := &RootOfTrie{}
	r.initTrie()
	r.insertIntoTrie("amogh")
	r.insertIntoTrie("amogh")
}
