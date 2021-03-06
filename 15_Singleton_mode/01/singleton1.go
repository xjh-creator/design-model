package main

import "fmt"

var s *singleton

type singleton struct {
}

func GetInstance() *singleton {
	if s == nil {
		s = &singleton{}
	}
	return s
}

func main() {
	a1 := GetInstance()
	a2 := GetInstance()
	fmt.Println(a1 == a2) //结果为true
}
