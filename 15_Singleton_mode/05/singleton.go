package main

import "fmt"

func init() {
	s = &singleton{}
}

var s *singleton

type singleton struct {
}

func GetInstance() *singleton {
	return s
}

func main() {
	a1 := GetInstance()
	a2 := GetInstance()
	fmt.Println(a1 == a2) //结果为true
}
