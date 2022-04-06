package main

import "fmt"

// 实现Component抽象类以及ConcreteComponent具体构建
type Component interface {
	show()
}

//Person 是ConcreteComponent具体构建
type Person struct {
    name string
}

func (p *Person)show()  {
	fmt.Println("装扮的",p.name)
}

//Finery 是 Decorator抽象装饰
type Finery interface {
	Component
}

//TShirts 是ConcreteDecorator，具体装饰
type TShirts struct {
	Finery
}

func (p *TShirts)show()  {
	fmt.Println("大T-shirt")
	p.Finery.show()
}

type BigTrouser struct {
	Finery
}

func (p *BigTrouser)show()  {
	fmt.Println("跨裤")
	p.Finery.show()
}

type BigSneakers struct {
	Finery
}

func (p *BigSneakers)show()  {
	fmt.Println("破球鞋")
	p.Finery.show()
}

type BigSuit struct {
	Finery
}

func (p *BigSuit)show()  {
	fmt.Println("西装")
	p.Finery.show()
}

type BigTie struct {
	Finery
}

func (p *BigTie)show()  {
	fmt.Println("领带")
	p.Finery.show()
}

type BigLeatherShoes struct {
	Finery
}

func (p *BigLeatherShoes)show()  {
	fmt.Println("皮鞋")
	p.Finery.show()
}

func  NewDecorator(t string,decorator Finery) Finery {
	switch t {
	case "TShirts":
		return &TShirts{decorator}
	case "BigTrouser":
		return &BigTrouser{decorator}
	case "BigSneakers":
		return &BigSneakers{decorator}
	case "BigSuit":
		return &BigSuit{decorator}
	case "BigTie":
		return &BigTie{decorator}
	case "BigLeatherShoes":
		return &BigLeatherShoes{decorator}
	default:
		return nil
	}
}

func main()  {
	component := &Person{"小明"}
	concreteDecorator1 := NewDecorator("TShirts",component)
	concreteDecorator1.show()

	concreteDecorator2 := NewDecorator("BigSuit",component)
	concreteDecorator2.show()
}
