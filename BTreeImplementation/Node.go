package main

import "log"

type Node struct {
	isLeaf bool
	nodes  []*Node
	parent *Node
	keys   []int
	size   int
}

func (nd *Node) PrintNode() {
	for _, k := range nd.keys {
		log.Println(k)
	}
}

func (nd *Node) GetNodes() []*Node {
	return nd.nodes
}

func (nd *Node) GetParent() *Node {
	return nd.parent
}

func (nd *Node) IsLeaf() bool { 
	return nd.isLeaf
}