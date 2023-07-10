package main

import (
	"log"
)

func main(){
	var firstSlice = make([]string,10)

	firstSlice[0] = "RandomString1"

	log.Printf(firstSlice[0])

	getSlice(firstSlice)
	
	log.Printf(firstSlice[1])
	getSliceByPointer(&firstSlice)

	d:= []byte{'n','i','k','u','l'}

	e:= d[:3]

	log.Println(d)
	log.Println(d[0])
	log.Println(e)
	log.Println(e[0])
	
	//

	var tS = NewTempStruct("a",2)

	tS.DoSomething()

}

func getSlice(slc []string){
	log.Println(slc)
	slc[1] = "RandomString2"
	log.Println(slc)
}

func getSliceByPointer(slp *[]string){
	log.Println(&slp)
	log.Println(slp)
	log.Println(*slp)
	
}