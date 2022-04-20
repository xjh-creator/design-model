package main

import "fmt"

type IHandsetBrand interface {
	SetHandsetSoft(soft IHandsetSoft)
	Run()
}

type HandsetBrandM struct {
	Soft IHandsetSoft
}

func NewHandsetBrandM(handsetSoft IHandsetSoft) *HandsetBrandM {
	return &HandsetBrandM{handsetSoft}
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

func NewHandsetBrandN(handsetSoft IHandsetSoft) *HandsetBrandN {
	return &HandsetBrandN{handsetSoft}
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

func NewIHandsetBrand(handsetBrand string,handsetSoft IHandsetSoft) IHandsetBrand {
	switch handsetBrand {
	case "M":return NewHandsetBrandM(handsetSoft)
	case "N":return NewHandsetBrandN(handsetSoft)
	default:
		return nil
	}
}

func main()  {
	handsetBrandM := NewIHandsetBrand("M",NewHandsetGame())
	handsetBrandM.Run()
	handsetBrandM.SetHandsetSoft(NewHandsetBrandAddressList())
	handsetBrandM.Run()

	handsetBrandN := NewIHandsetBrand("N",NewHandsetGame())
	handsetBrandN.Run()
	handsetBrandN.SetHandsetSoft(NewHandsetBrandAddressList())
	handsetBrandN.Run()
}