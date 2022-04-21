package main

import (
	"fmt"
	"time"
)

type Command interface {
	ExcuteCommand()
	BakeName() string
}

type BakeMuttonCommand struct {
	name string
	receiver *Barbecuer
}

func NewBakeMuttonCommand(receiver *Barbecuer) *BakeMuttonCommand {
	return &BakeMuttonCommand{name:"烤羊肉串",receiver: receiver}
}

func (b BakeMuttonCommand)BakeName() string {
	return b.name
}

func (b BakeMuttonCommand)ExcuteCommand()  {
	b.receiver.BakeMutton()
}

type BakeChickenWingCommand struct {
	name string
	receiver *Barbecuer
}

func NewBakeChickenWingCommand(receiver *Barbecuer) *BakeChickenWingCommand {
	return &BakeChickenWingCommand{name:"烤鸡翅",receiver: receiver}
}

func (b BakeChickenWingCommand)BakeName() string {
	return b.name
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
	orders []Command
}

func NewWaiter() *Waiter {
	return &Waiter{}
}

func (w *Waiter)SetOrder(command Command)  {
	w.orders = append(w.orders,command)
	fmt.Println("增加订单：",command.BakeName()," 时间：",time.Now().Format("2006-01-01 11:32:31"))
}

func (w *Waiter)CancelOrder(command Command)  {
	for i,v := range w.orders{
		if v == command{
			w.orders = append(w.orders[:i],w.orders[i+1:]...)
			break
		}
	}
	fmt.Println("删除订单：",command.BakeName()," 时间：",time.Now().Format("2006-01-01 11:32:31"))
}

func (w *Waiter)Notify()  {
	for _,v := range w.orders{
		v.ExcuteCommand()
	}
}

func main()  {
	//开店前的准备
	boy := NewBarbecuer()
	bakeMuttonCommand1 := NewBakeMuttonCommand(&boy)
	bakeMuttonCommand2 := NewBakeMuttonCommand(&boy)
	bakeChickenWingCommand1 := NewBakeChickenWingCommand(&boy)
	girl := NewWaiter()

	//开门营业,顾客点菜
	girl.SetOrder(bakeMuttonCommand1)
	girl.SetOrder(bakeMuttonCommand2)
	girl.SetOrder(bakeChickenWingCommand1)
    girl.CancelOrder(bakeMuttonCommand1)
	//点菜完毕，通知厨房
	girl.Notify()
}


