package main

import "fmt"

type IHandsetBrand interface {
	SetHandsetSoft(soft IHandsetSoft)
	Run()
}

type HandsetBrandM struct {
	Soft IHandsetSoft
}

func NewHandsetBrandM() *HandsetBrandM {
	return &HandsetBrandM{nil}
}

func (h *HandsetBrandM)SetHandsetSoft(soft IHandsetSoft)  {
	h.Soft = soft
}

func (h *HandsetBrandM)Run()  {
	fmt.Printf("手机品牌M")
	h.Soft.Run()
}

type HandsetBrandN struct {
	Soft IHandsetSoft
}

func NewHandsetBrandN() *HandsetBrandN {
	return &HandsetBrandN{nil}
}

func (h *HandsetBrandN)SetHandsetSoft(soft IHandsetSoft)  {
	h.Soft = soft
}

func (h *HandsetBrandN)Run()  {
	fmt.Printf("手机品牌N")
	h.Soft.Run()
}


type IHandsetSoft interface {
	Run()
}

type HandsetGame struct {
}

func NewHandsetGame() *HandsetGame {
	return &HandsetGame{}
}

func (h HandsetGame)Run(){
	fmt.Println("运行手机游戏")
}

type HandsetBrandAddressList struct {
	*HandsetBrandN
}

func NewHandsetBrandAddressList() *HandsetBrandAddressList {
	return &HandsetBrandAddressList{}
}

func (h HandsetBrandAddressList)Run(){
	fmt.Println("运行手机通信录")
}

func main()  {
	var ab IHandsetBrand
	ab = NewHandsetBrandN()

	ab.SetHandsetSoft(NewHandsetGame())
	ab.Run()

	ab.SetHandsetSoft(NewHandsetBrandAddressList())
	ab.Run()

	ab = NewHandsetBrandM()

	ab.SetHandsetSoft(NewHandsetGame())
	ab.Run()

	ab.SetHandsetSoft(NewHandsetBrandAddressList())
	ab.Run()

}