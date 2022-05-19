package main

import (
	"sync"
)

var (
	s    *singleton
	lock = &sync.Mutex{}
)

type singleton struct {
}

func GetInstance() *singleton {
	if s == nil {
		lock.Lock()
		defer lock.Unlock()
		if s == nil {
			s = &singleton{}
		}
	}
	return s
}
