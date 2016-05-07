package ds

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	BTREE_PRIME = 71
)

type BTreeNode struct {
	Depth int `json:"depth"` // the number of descendent nodes

	Parent *BTreeNode `json:"-"`
	Left   *BTreeNode `json:"left"`
	Right  *BTreeNode `json:"right"`

	Key   *int   `json:"key"`
	Value string `json:"value"` // this could be json data or whatever
}

type BTree struct {
	Root *BTreeNode `json:"root"`
}

//==================================
// Public methods
//==================================
func (b *BTree) Put(key string, value string) {
	k := b.key(key)

	if b.Root == nil {
		b.Root = &BTreeNode{
			Parent: nil,
			Key:    &k,
			Value:  value,
		}
		return
	}

	b.Root.insert(k, value)
}

func (b *BTree) Get(key string) {

}

func (b *BTree) Del(key string) {

}

func (b *BTree) Debug() {
	//j, _ := json.Marshal(b)
	j, _ := json.MarshalIndent(b, "", "\t")
	fmt.Println(string(j))
}

//==================================
// Private methods
//==================================
func (b *BTreeNode) insert(key int, value string) {
	if key < *b.Key {
		b.Depth++

		if b.Left == nil {
			b.Left = &BTreeNode{
				//Parent: b,
				Key:   &key,
				Value: value,
			}
			return
		}

		b.Left.insert(key, value)

	} else {
		b.Depth--

		if b.Right == nil {
			b.Right = &BTreeNode{
				//Parent: b,
				Key:   &key,
				Value: value,
			}
			return
		}

		b.Right.insert(key, value)
	}

	b.balance()
}

func (b *BTreeNode) balance() {
	if b.Depth < -1 {
		b.Depth++
		b.rotate(true)
	} else if b.Depth > 1 {
		b.Depth--
		b.rotate(false)
	}
}

func (b *BTreeNode) rotate(rot_right bool) {
	n := *b

	if rot_right {
		var left *BTreeNode

		if b.Left != nil {
			left = b.Left
		}

		b.Key = b.Right.Key
		b.Value = b.Right.Value
		b.Left = b.Right
		b.Right = b.Right.Right
		b.Left.Right = b.Left.Left

		if left != nil {
			b.Left.Left = left
		}

		b.Left.Key = n.Key
		b.Left.Value = n.Value
		return
	}

	var right *BTreeNode

	if b.Right != nil {
		right = b.Right
	}

	b.Key = b.Left.Key
	b.Value = b.Left.Value
	b.Right = b.Left
	b.Left = b.Left.Left
	b.Right.Left = b.Right.Right

	// there's some bug here...
	if right != nil {
		b.Right.Right = right
	}

	// this is a hacky fix for that bug..
	if b.Right.Left == b.Right.Right {
		b.Right.Left = nil
	}

	b.Right.Key = n.Key
	b.Right.Value = n.Value
}

func (BTree) key(key string) int {
	k, err := strconv.Atoi(key)

	if err == nil {
		return k
	}

	b := []byte(key)

	for i, _ := range b {
		k += int(b[i])
	}
	return k % BTREE_PRIME
}
