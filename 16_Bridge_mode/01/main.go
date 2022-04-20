package main

import "fmt"

type IHandsetBrand interface {
	Run()
}

type HandsetBrandM struct {
	Name string
}

func NewHandsetBrandM() *HandsetBrandM {
	return &HandsetBrandM{
		Name: "M品牌",
	}
}

type HandsetBrandN struct {
    Name string
}

func NewHandsetBrandN() *HandsetBrandN {
	return &HandsetBrandN{
		Name: "N品牌",
	}
}

type HandsetBrandMGame struct {
	*HandsetBrandM
}

func (h HandsetBrandMGame)Run(){
	fmt.Println("运行",h.Name,"手机游戏")
}

type HandsetBrandMAddressList struct {
	*HandsetBrandM
}

func (h HandsetBrandMAddressList)Run(){
	fmt.Println("运行",h.Name,"手机通信录")
}

type HandsetBrandNGame struct {
	*HandsetBrandN
}

func (h HandsetBrandNGame)Run(){
	fmt.Println("运行",h.Name,"手机游戏")
}

type HandsetBrandNAddressList struct {
	*HandsetBrandN
}

func (h HandsetBrandNAddressList)Run(){
	fmt.Println("运行",h.Name,"手机通信录")
}

func main()  {
	var ab IHandsetBrand
	ab = HandsetBrandMAddressList{NewHandsetBrandM()}
	ab.Run()

	ab = HandsetBrandMGame{NewHandsetBrandM(),}
	ab.Run()

	ab = HandsetBrandNAddressList{NewHandsetBrandN()}
	ab.Run()

	ab = HandsetBrandNGame{NewHandsetBrandN()}
	ab.Run()
}
