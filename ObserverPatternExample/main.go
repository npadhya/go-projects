package main

import (
	"log"
	"runtime"

	"ObserverPatternExample/concrete"
)

func main() {
	pc, _, _, _ := runtime.Caller(0)
	methodName := runtime.FuncForPC(pc).Name()
	log.Println(methodName, " Started")

	s := &concrete.ConcreteSubject{}

	o1 := &concrete.ConcreteObserver{ID: 1}
	o2 := &concrete.ConcreteObserver{ID: 2}
	o3 := &concrete.ConcreteObserver{ID: 3}

	s.Register(o1)
	s.Register(o2)

	s.Notify()
	s.Deregister(o1)
	s.Deregister(o3)

	s.Notify()

	log.Println(methodName, " Ended")
}
