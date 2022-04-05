package main

import "fmt"

//实现一个简单的穿搭demo
//如果要增加装扮，就需要改写person类
//这种写法违背了开发-封闭原则：软件实体（类、函数、模块）可以扩展，但不可以修改


type Person struct {
	name string
}

func (p *Person)WearTShirts()  {
	fmt.Println("大T-shirt")
}

func (p *Person)WearBigTrouser()  {
	fmt.Println("跨裤")
}

func (p *Person)WearSneakers()  {
	fmt.Println("破球鞋")
}

func (p *Person)WearSuit()  {
	fmt.Println("西装")
}

func (p *Person)WearTie()  {
	fmt.Println("领带")
}

func (p *Person)WearLeatherShoes()  {
	fmt.Println("皮鞋")
}

func (p *Person)show()  {
	fmt.Println("装扮的",p.name)
}

func main()  {
	person := &Person{"小菜"}
	fmt.Println("第一种装扮:")
	person.WearTShirts()
	person.WearBigTrouser()
	person.WearSneakers()
	person.show()

	fmt.Println("第二种装扮:")
	person.WearSuit()
	person.WearTie()
	person.WearLeatherShoes()
	person.show()

}
