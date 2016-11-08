package ds

type LLNode struct {
	Data interface{}
	Next *LLNode
}

type LinkedList struct {
	// Root origin node
	Root *LLNode

	// Prev keeps track of the previously inserted node
	Prev *LLNode

	// Last keeps track of the latest node
	Last *LLNode
}

// Insert appends a node to list
func (list *LinkedList) Insert(node *LLNode) {
	if list.Root == nil {
		list.Root = node
		list.Prev = list.Root
		list.Last = list.Prev.Next
		return
	}

	for {
		if list.Last == nil {
			list.Last = node
			list.Prev.Next = node
			break
		}
		list.Prev = list.Last
		list.Last = list.Prev.Next
	}
}

// InsertAt appends a node at a given location based on the location of another node.
func (list *LinkedList) InsertAt(newNode *LLNode, pivotNode *LLNode, before bool) {
	if before == true {
		list.InsertBefore(newNode, pivotNode)
	} else {
		newNode.Next = pivotNode.Next
		pivotNode.Next = newNode
	}
}

// InsertBefore insert the new node where the given node currently is and makes the current node
// the next node in the chain of the new node.
func (list *LinkedList) InsertBefore(newNode *LLNode, pivotNode *LLNode) {
	if list.Root == pivotNode {
		tmp := list.Root
		list.Root = newNode
		list.Root.Next = tmp
		return
	}

	prev := list.Root

	for {
		if prev.Next == pivotNode {
			newNode.Next = prev.Next
			prev.Next = newNode
			break
		}
		prev = prev.Next
	}
}

func (list *LinkedList) Remove(node *LLNode) bool {
	if node == nil {
		return false
	}

	current := list.Root

	// we track the current pointer so that we can overwrite it
	// when needed.
	prev := list.Root

	for {
		if current == nil {
			return false
		}

		if current == node {
			// if the node to be removed is the last...
			if list.Last == node {
				list.Last = nil
				list.Prev.Next = nil
				return true
			}

			// if the node to be removed is the root...
			if current == list.Root {
				tmp := current.Next
				list.Root = tmp
				current = tmp
			}

			// shift node pointers down
			if current.Next != nil {
				prev.Next = current.Next
			}

			return true
		}

		// update the previous
		prev = current
		// update the current
		current = current.Next
	}
}

func (list *LinkedList) Find(node *LLNode) *LLNode {
	current := list.Root

	for {
		if current == nil {
			return nil
		}

		if current == node {
			return current
		}
		current = current.Next
	}
}

func (list *LinkedList) FindByValue(value interface{}) *LLNode {
	current := list.Root

	for {
		if current == nil {
			return nil
		}

		if current.Data == value {
			return current
		}
		current = current.Next
	}
}
