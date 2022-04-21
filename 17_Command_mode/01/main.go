package main

import "fmt"

//Barbecuer烤肉串者
type Barbecuer struct {
	
}

func NewBarbecuer() Barbecuer {
	return Barbecuer{}
}

func (b Barbecuer)BakeMutton()  {
	fmt.Println("烤羊肉串！")
}

func (b Barbecuer)BakeChickenWing()  {
	fmt.Println("烤鸡翅！")
}

func main()  {
	boy := NewBarbecuer()
	boy.BakeMutton()
	boy.BakeMutton()
	boy.BakeMutton()
	boy.BakeChickenWing()
	boy.BakeMutton()
	boy.BakeMutton()
	boy.BakeChickenWing()
}


