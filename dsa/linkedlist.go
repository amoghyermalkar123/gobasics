package dsa

type ListNode struct {
	node *ListNode
	word string
	// isLast bool
}

func InsertNode(node *ListNode, word string) {
	if node.word == "" {
		node.word = word
		node.node = &ListNode{}
	} else {
		currentNode := &node.node
		InsertNode(*currentNode, word)
	}
}

// func StartList() {
// 	ln := &ListNode{}
// 	InsertNode(ln, "amogh")
// 	InsertNode(ln, "ishan")
// 	fmt.Println(ln.node)
// }
