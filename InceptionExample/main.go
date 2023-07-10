package main

import (
	"log"

	"golang.org/x/exp/constraints"
)

type Dream[REALITY constraints.Ordered] func(REALITY) func(string) string

func inception[REALITY constraints.Ordered] (reality REALITY) Dream[REALITY] {
	log.Println(reality)
	return Dream[REALITY](func(dream REALITY) func(s string) string {
		log.Println(dream)
		return func(s string) string {
			log.Println(s)
			return "Returning"
		}
	})

}

func logging(x func()) func(){
	x()
	j := func(){
		x()
		log.Println("In side logging")
	}
	return j
}

func main() {
	l := inception("Level1")("Level2")("Level3")

	k := logging(func() {
		log.Println("Printing in logging")
	})
	k()
	log.Println(l)
	log.Println(&k)
	
}
