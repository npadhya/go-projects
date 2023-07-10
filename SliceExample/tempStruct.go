package main

import "log"

type tempStruct struct {
	name string
	num  int
}

func NewTempStruct(name string, num int) *tempStruct {
	return &tempStruct{
		name: name,
		num:  num,
	}
}

func (ts *tempStruct) DoSomething() {
	log.Println("Doing something")
	log.Println(ts.name)
}
