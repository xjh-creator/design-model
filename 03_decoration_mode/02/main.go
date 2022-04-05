package main

import "fmt"

type Person struct {
	name string
}

func (p *Person)show()  {
	fmt.Println("装扮的",p.name)
}


type Finery interface {
	show()
}

type TShirts struct {
}

func (p *TShirts)show()  {
	fmt.Println("大T-shirt")
}

type BigTrouser struct {
}

func (p *BigTrouser)show()  {
	fmt.Println("跨裤")
}

type BigSneakers struct {
}

func (p *BigSneakers)show()  {
	fmt.Println("破球鞋")
}

type BigSuit struct {
}

func (p *BigSuit)show()  {
	fmt.Println("西装")
}

type BigTie struct {
}

func (p *BigTie)show()  {
	fmt.Println("领带")
}

type BigLeatherShoes struct {
}

func (p *BigLeatherShoes)show()  {
	fmt.Println("皮鞋")
}

//相比与01，优化的点是运用了面向对象的思想，避免违反了开发-封闭原则
//但这里还存在问题，每次显示，都要调用一次show,这种方式很不优雅
func main()  {
	person := &Person{name: "xiaoming",}
	fmt.Println("第一种装扮:")
	person.show()
	tshirt := new(TShirts)
	tshirt.show()
	bigTie := new(BigTie)
	bigTie.show()
}

