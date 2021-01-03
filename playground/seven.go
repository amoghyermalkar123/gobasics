package playground

import "fmt"

// hash table with linked list storage to avoid collision

// linked list insertion : insert  at the start of the list i.e. update the root node at every insert

const maxSize = 8

type RootNodee struct {
	root *Nodee
}

type Nodee struct {
	word string
	next *Nodee
}

type HTable struct {
	table [maxSize]*RootNodee
}

func (ht *HTable) init() {
	for i := range ht.table {
		ht.table[i] = &RootNodee{root: &Nodee{}}
	}
	// fmt.Println("init table:", ht.table)
}
func (ht *HTable) insertNode(word string) {
	hashV := hash(word)

	if ht.search(hashV, word) == false {
		// word not found..hence insert
		node := &Nodee{
			word: word,
			next: ht.table[hashV].root,
		}
		ht.table[hashV].root = node
		fmt.Println("word inserted")
	} else {
		fmt.Println("word already exists in the hash table")
	}

}

func (ht *HTable) search(hash int, word string) bool {
	currentNode := ht.table[hash].root
	for currentNode.next != nil {
		if currentNode.word == word {
			return true
		} else {
			currentNode = currentNode.next
		}
	}
	return false
}
func hash(word string) (hashValue int) {
	sum := 0
	for i := range word {
		sum += int(word[i] - 'a')
	}
	hashValue = sum % maxSize
	return hashValue
}

func StartHashI() {
	ht := HTable{}
	ht.init()
	ht.insertNode("amogh")
	ht.insertNode("ishan")
	ht.insertNode("nahsi")
}
