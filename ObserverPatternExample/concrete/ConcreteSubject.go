package concrete

import (
	"log"
	"runtime"

	"ObserverPatternExample/interfaces"
)

type ConcreteSubject struct {
	this_observers []interfaces.Observer
}

func (concretesub *ConcreteSubject) Register(observer interfaces.Observer) {
	pc, _, _, _ := runtime.Caller(0)
	methodName := runtime.FuncForPC(pc).Name()
	log.Println(methodName, " Started")
	concretesub.this_observers = append(concretesub.this_observers, observer)
	log.Println(methodName, " Ended")
}

func (concretesub *ConcreteSubject) Deregister(observer interfaces.Observer) {
	pc, _, _, _ := runtime.Caller(0)
	methodName := runtime.FuncForPC(pc).Name()
	log.Println(methodName, " Started")
	for i, obs := range concretesub.this_observers {
		if obs == observer {
			log.Println("Doing deregisteration", observer)
			concretesub.this_observers = append(concretesub.this_observers[:i], concretesub.this_observers[i+1:]...)
			log.Println(methodName, " Ended")
			return
		}
	}
	log.Println("No such Observer found", observer)
	log.Println(methodName, " Ended")
}

func (concretesub *ConcreteSubject) Notify() {
	pc, _, _, _ := runtime.Caller(0)
	methodName := runtime.FuncForPC(pc).Name()
	log.Println(methodName, " Started")
	for _, observer := range concretesub.this_observers {
		observer.Update()
	}
	log.Println(methodName, " Ended")
}
