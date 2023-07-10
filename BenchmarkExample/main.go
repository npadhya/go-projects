package main

import (
	"log"
	"reflect"
	"runtime"
)

type Student struct {
	name string `json:"name"`
}

func (*Student) GetStudent() {
	log.Println("v ...any")
}

func catch() {
	pc, _, _, _ := runtime.Caller(0)
	log.SetPrefix(runtime.FuncForPC(pc).Name() + " | ")
	log.Println("catch called")
	if r := recover(); r != nil {
		log.Println(reflect.TypeOf(r))
		switch t := r.(type) {
		case string:
			log.Println("string", t)
		case Student:
			log.Println("Student", t.name)
		}
		log.Println("Recovering from ", r)
	}
}

func doSomething() {
	pc, _, _, _ := runtime.Caller(0)
	log.SetPrefix(runtime.FuncForPC(pc).Name() + " | ")
	defer catch()
	// do something that might panic
	log.Println("Before something went wrong in doSomething")
	log.Println(runtime.FuncForPC(pc).Name())
	panic(Student{"Error"})
}

func doSomething2() {
	defer catch()
	pc, _, _, _ := runtime.Caller(0)
	log.SetPrefix(runtime.FuncForPC(pc).Name() + " | ")
	// do something that might panic
	log.Println("Before something went wrong in doSomething2")
	log.Println(runtime.FuncForPC(pc).Name())
	nested()
	panic("Error")
}

func nested() {
	pc, _, _, _ := runtime.Caller(0)
	log.Println(runtime.FuncForPC(pc).Name())
}

func main() {
	pc, _, _, _ := runtime.Caller(0)
	log.SetPrefix(runtime.FuncForPC(pc).Name() + " | ")
	log.Println(runtime.FuncForPC(pc).Name())
	log.Println("--", runtime.FuncForPC(pc))
	doSomething()
	doSomething2()
	log.Println("This statement will run after panic")
	log.Println("++", runtime.FuncForPC(pc))
	// os.Exit(0)
}
