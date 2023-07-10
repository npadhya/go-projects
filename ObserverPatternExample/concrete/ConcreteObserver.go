package concrete

import (
	"log"
	"runtime"
)

type ConcreteObserver struct {
	ID int
}

func (concreteObserver *ConcreteObserver) Update() {
	pc, _, _, _ := runtime.Caller(0)
	methodName := runtime.FuncForPC(pc).Name()
	log.Println(methodName, " Started")
	log.Println("Observer : ", concreteObserver.ID, " updated")
	log.Println(methodName, " Ended")
}
