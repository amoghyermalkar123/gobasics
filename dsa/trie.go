package dsa

import "fmt"

const max = 26

type Node struct {
	children [max]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func initiateRoot() *Trie {
	// here initiatedRoot stores address and if you want to return it,
	// the return type is a *value because you return a read-thru-address variable
	// when you want to return an address
	initiatedRoot := &Trie{root: &Node{}}
	return initiatedRoot
}

func Insert(t *Trie, word string) (trie *Trie) {
	currentNode := t.root

	for i, _ := range word {
		characterIndex := word[i] - 'a'
		if currentNode.children[characterIndex] == nil {
			currentNode.children[characterIndex] = &Node{}
		}
		currentNode = currentNode.children[characterIndex]
	}
	currentNode.isEnd = true
	trie = t
	return trie
}

func Search(t *Trie, word string) (ans bool) {
	currentNode := t.root
	for i, _ := range word {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		} else {
			currentNode = currentNode.children[charIndex]
		}
	}
	if currentNode.isEnd == true {
		return true
	} else {
		return false
	}
}

func Start() {
	root := initiateRoot()
	// root above is a read-thru-address variable which you can pass to Insert() function
	_ = Insert(root, "amogh")
	answer := Search(root, "amoghh")
	fmt.Println(answer)
}
