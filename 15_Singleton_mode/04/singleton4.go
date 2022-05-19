package main

import (
	"sync"
)

var (
	s    *singleton
	once = &sync.Once{}
)

type singleton struct {
}

func GetInstance() *singleton {
	once.Do(func() {
		s = &singleton{}
	})
	return s
}
