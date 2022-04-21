package main

import "fmt"

type Command interface {
	ExcuteCommand()
}

type BakeMuttonCommand struct {
	receiver *Barbecuer
}

func NewBakeMuttonCommand(receiver *Barbecuer) *BakeMuttonCommand {
	return &BakeMuttonCommand{receiver: receiver}
}

func (b BakeMuttonCommand)ExcuteCommand()  {
	b.receiver.BakeMutton()
}

type BakeChickenWingCommand struct {
	receiver *Barbecuer
}

func NewBakeChickenWingCommand(receiver *Barbecuer) *BakeChickenWingCommand {
	return &BakeChickenWingCommand{receiver: receiver}
}

func (b BakeChickenWingCommand)ExcuteCommand()  {
	b.receiver.BakeChickenWing()
}

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

type Waiter struct {
	command Command
}

func NewWaiter() *Waiter {
	return &Waiter{}
}

func (w *Waiter)SetOrder(command Command)  {
	w.command = command
}

func (w *Waiter)Notify()  {
	w.command.ExcuteCommand()
}

func main()  {
	//开店前的准备
	boy := NewBarbecuer()
	bakeMuttonCommand1 := NewBakeMuttonCommand(&boy)
	bakeMuttonCommand2 := NewBakeMuttonCommand(&boy)
	bakeChickenWingCommand1 := NewBakeChickenWingCommand(&boy)
	girl := NewWaiter()

	//开门营业
	girl.SetOrder(bakeMuttonCommand1)
	girl.Notify()
	girl.SetOrder(bakeMuttonCommand2)
	girl.Notify()
	girl.SetOrder(bakeChickenWingCommand1)
	girl.Notify()
}


