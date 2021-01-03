package dsa

import "fmt"

const ArraySize = 7

type Hashtable struct {
	table [ArraySize]*HeadNode
}

type HeadNode struct {
	rootNode *Nodes
}

type Nodes struct {
	word     string
	nextNode *Nodes
}

func CreateHashTable(table *Hashtable, word string) {
	index := 0
	for i := range word {
		characterIndex := word[i] - 'a'
		index += int(characterIndex)
		index = index % ArraySize
	}

	// you cant do the below operation
	// table.table[index].rootNode.word = "word"
	// so do this =>
	table.table[index].ListInsert(word)

}

func (hn *HeadNode) ListInsert(word string) {
	// NOTE:
	// still cant do this
	// hn.rootNode.word = word

	// doc starts

	// search if the key already exists
	if !hn.ListSearch(word) {
		node := &Nodes{word: word}
		// inserting our newest node on the top and
		// making it a head after that
		node.nextNode = hn.rootNode
		hn.rootNode = node
	} else {
		fmt.Println(word, "already exists")
	}
}

func (ht *Hashtable) SearchHashTable(word string) {
	index := 0
	for i := range word {
		characterIndex := word[i] - 'a'
		index += int(characterIndex)
		index = index % ArraySize
	}

	if ht.table[index].ListSearch(word) == true {
		fmt.Println("word found")
	} else {
		fmt.Println("word not found")
	}
}

func (hn *HeadNode) ListSearch(word string) bool {
	currentNode := hn.rootNode
	for currentNode != nil {
		if currentNode.word == word {
			return true
		} else {
			currentNode = currentNode.nextNode
		}
	}
	return false
}

func Init() *Hashtable {
	res := &Hashtable{}
	for i := range res.table {
		res.table[i] = &HeadNode{}
	}
	return res
}

func Start2() {
	table := Init()
	CreateHashTable(table, "amogh")
	// fmt.Println(table.table[4].rootNode.word)
	table.SearchHashTable("amogh")
}
