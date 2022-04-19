package main

import "fmt"

var singleton *Singleton

type Singleton struct {

}

func GetInstance() *Singleton {
	if singleton == nil{
		singleton = &Singleton{}
	}
	return singleton
}

func main()  {
	a1 := GetInstance()
	a2 := GetInstance()
	fmt.Println(a1 == a2) //结果为true
}
