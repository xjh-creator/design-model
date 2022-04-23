package main

import "fmt"

type Mediator interface {
	Send(message string,colleague Colleague)
}

type ConcreteMediator struct {
	colleague1 Colleague
	colleague2 Colleague
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{}
}

func (c *ConcreteMediator)Send(message string,colleague Colleague)  {
	if c.colleague1 == colleague{
		c.colleague2.Notify(message)
	}else{
		c.colleague1.Notify(message)
	}
}

type Colleague interface {
	Notify(message string)
	Send(message string)
}

type ConcreteColleague1 struct {
	mediator Mediator
}

func NewConcreteColleague1(mediator Mediator) *ConcreteColleague1 {
	return &ConcreteColleague1{mediator: mediator}
}

func (c *ConcreteColleague1)Send(message string)  {
	c.mediator.Send(message,c)
}

func (c *ConcreteColleague1)Notify(message string)  {
	fmt.Println("同事1得到消息",message)
}

type ConcreteColleague2 struct {
	mediator Mediator
}

func NewConcreteColleague2(mediator Mediator) *ConcreteColleague2 {
	return &ConcreteColleague2{mediator: mediator}
}

func (c *ConcreteColleague2)Send(message string)  {
	c.mediator.Send(message,c)
}

func (c *ConcreteColleague2)Notify(message string)  {
	fmt.Println("同事2得到消息",message)
}

func main()  {
	m := NewConcreteMediator()

	c1 := NewConcreteColleague1(m)
	c2 := NewConcreteColleague2(m)

	m.colleague1 = c1
	m.colleague2 = c2

	c1.Send("吃饭了吗？")
	c2.Send("没有呢，你打算请客？")
}
