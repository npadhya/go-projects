package main

import "log"

var (
	MAX_KEYS     int
)

type BTree struct {
	this_root    *Node
	MAX_KEYS     int
}

func (btree *BTree) GetBTreeRoot() *Node {
	return btree.this_root
}

func (btree *BTree) PrintTree(nd *Node) {
	nd.PrintNode()

	for _, eachChild := range nd.GetNodes() {
		btree.PrintTree(eachChild)
	}
}

func (btree *BTree) Find(key int) *Node {
	if btree.this_root == nil {
		return nil
	} else {
		current := btree.this_root

		for !current.isLeaf {
			for i := 0; i < current.size; i++ {
				if key < current.keys[i] {
					current = current.nodes[i]
					break
				}

				if i == current.size-1 {
					current = current.nodes[i+1]
					break
				}
			}

			for j := 0; j < current.size; j++ {
				if current.keys[j] == key {
					return current
				}
			}
		}
		return nil
	}
}

func (btree *BTree) FindParent(node *Node) *Node {
	return node.GetParent()
}

func (btree *BTree) Insert(key int) {
	log.Println("Inserting ", key, " Node into the tree")
	if btree.this_root == nil {
		n := &Node{isLeaf : true, keys : make([]int, MAX_KEYS), nodes: make([]*Node, MAX_KEYS+1), size : 1 }
		n.keys[0] = key
		btree.this_root = n
	} else {
		current := btree.this_root
		parent := current.GetParent()

		for !current.IsLeaf() {

			parent = current

			for i , k := range current.keys {
				if key < k {
					current = current.nodes[i]
					break
				}

				if i == current.size - 1{
					current = current.nodes[i+1]
					break
				}
			}
		}

		if current.size < MAX_KEYS {
			i :=0
			for key > current.keys[i] && i < current.size{
				i++
			}

			for j , _ := range current.keys {
				current.keys[j] = current.keys[j - 1]
			}

			current.keys[i] = key

			current.size++

			current.nodes[current.size] = current.nodes[current.size - 1]

			current.nodes[current.size - 1] = nil
		} else {
			leafNode := &Node{isLeaf : true, keys : make([]int, MAX_KEYS), nodes: make([]*Node, MAX_KEYS+1), size : 1 }
			tempNode := make([]int,MAX_KEYS + 1)
			
			for i , v := range current.keys {
				tempNode[i] = v
			}

			var i int
			var j int

			for key > tempNode[i] && i < MAX_KEYS {
				i++
			}

			for j = MAX_KEYS + 1; j > i ; j-- {
				tempNode[j] = tempNode[j - 1]
			}

			tempNode[i] = key

			leafNode.isLeaf = true
			current.size = (MAX_KEYS + 1) / 2
			leafNode.size = (MAX_KEYS + 1) - ((MAX_KEYS + 1) / 2)

			current.nodes[current.size] = leafNode

			leafNode.nodes[leafNode.size] = current.nodes[MAX_KEYS]

		}
	}
}

func (btree *BTree) DeleteNode(key int) {
	log.Println("Deleting ", key, " Node from the tree")
}

func (btree *BTree) ShiftLevel(key int, a *Node, b *Node) {
	log.Println("Shifting the tree")
}
