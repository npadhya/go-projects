package main

import "log"

type Block struct {
	ID   string
	data *MainData
}

func (blk *Block) PrintBlock() {
	log.Println(blk.ID)
	blk.data.PrintMainData()
}
