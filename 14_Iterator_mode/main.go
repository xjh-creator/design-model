package main

import "fmt"

type Iterator interface {
	First() string
	Next() string
	IsDone() bool
	CurrentItem() string
}

type ConcreteIterator struct {
	*ConcreteAggregate
	current int
}

func NewConcreteIterator(aggregate *ConcreteAggregate) *ConcreteIterator {
	return &ConcreteIterator{aggregate,0}
}

func (c ConcreteIterator)First() string {
	return c.Items[0]
}

func (c *ConcreteIterator)Next() string {
	c.current++
	if c.current < len(c.Items){
		return c.Items[c.current]
	}
	return ""
}

func (c ConcreteIterator)IsDone() bool {
	return c.current >= len(c.Items)
}

func (c ConcreteIterator)CurrentItem() string {
	return c.Items[c.current]
}

type Aggregate interface {
	CreateIterator() Iterator
}

type ConcreteAggregate struct {
	Items []string
	Count int
	index int
}

func NewConcreteAggregate() *ConcreteAggregate {
	return &ConcreteAggregate{
		make([]string, 0),
		0,
		0,
	}
}

func (c *ConcreteAggregate)CreateIterator() Iterator {
	return &ConcreteIterator{c,0}
}

func main()  {
	bus := NewConcreteAggregate() //公交车，聚集对象
	bus.Items = append(bus.Items,"大鸟","小菜","行李","老外","公交内部员工","小偷")

	seller := NewConcreteIterator(bus) //售票员，迭代器对象
	for !seller.IsDone(){
		fmt.Println(seller.CurrentItem(),"请买票！")
		seller.Next()
	}
}


