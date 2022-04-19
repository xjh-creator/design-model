package main

import (
	"sync"
)

var singleton *Singleton
var lock *sync.Mutex = &sync.Mutex{}

type Singleton struct {

}

func GetInstance() *Singleton {
	if singleton == nil{
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil{
			singleton = &Singleton{}
		}
	}
	return singleton
}


