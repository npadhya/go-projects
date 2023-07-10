package main

import "log"

type MainData struct {
	Name    string
	Age     string
	Address string
}

func (md *MainData) PrintMainData() {
	log.Println(md.Name, md.Age, md.Address)
}
