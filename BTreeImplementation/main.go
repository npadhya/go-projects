package main

func main() {
	bt := BTree{MAX_KEYS: 3}
	bt.Insert(1)
	bt.Insert(2)
	bt.Insert(3)
	bt.Insert(4)
	bt.Insert(5)
	bt.Insert(6)
	bt.Insert(7)
	bt.Insert(8)
	bt.Insert(9)
	bt.Insert(10)
	bt.Insert(11)

	root := bt.GetBTreeRoot()

	bt.PrintTree(root)
}
