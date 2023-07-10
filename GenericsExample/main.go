package main

import (
	"log"
	"golang.org/x/exp/constraints"
)


type CustomData interface {
	constraints.Ordered | ~string
}

type Student[T CustomData] struct{
	Data T
}

type LogAndReturnDefaultFunc[T CustomData] func(T) T

func iWillReturnAFunc[T CustomData](name T) LogAndReturnDefaultFunc[T] {
	log.Println(name)
	// or do something with the *name"
	fn := LogAndReturnDefaultFunc[T] (func(input T) T {
        t := name
		log.Println(input)
		// Do something with input
        return t
    })
    return fn
}

func Add[T CustomData](a T, b T) T {
	// Call iWillReturnAFunc, that will return a func
    returnedFunc := iWillReturnAFunc(a)
	// Run the returned func to get its o/p
	outputOfReturnedFunc := returnedFunc(b)
	//Do something with the Output of the func
	log.Println(outputOfReturnedFunc)
    return a
}

func main() {
	result := Add("abc", "DEF")
	result2 := Add(1.2, 3.4)

	// s1 := Student[string]{"Nikul"}
	// s2 := Student[string]{"Nikstring"}
	// res := Add(s1,s2)
	
	log.Printf("%v\n", result)
	log.Printf("%v\n", result2)
	// log.Println(res)
}