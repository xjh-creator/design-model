package main

import "fmt"

type UnitedNations interface {
	Declare(message string,country Country)
}

type UnitedNationsSecurityCouncil struct {
	colleague1 Country
	colleague2 Country
}

func NewUnitedNationsSecurityCouncil() *UnitedNationsSecurityCouncil {
	return &UnitedNationsSecurityCouncil{}
}

func (c *UnitedNationsSecurityCouncil)Declare(message string,country Country)  {
	if c.colleague1 == country{
		c.colleague2.GetMessage(message)
	}else{
		c.colleague1.GetMessage(message)
	}
}

type Country interface {
	Declare(message string)
	GetMessage(message string)
}

type USA struct {
	mediator UnitedNations
}

func NewUSA(mediator UnitedNations) *USA {
	return &USA{mediator: mediator}
}

func (c *USA)Declare(message string)  {
	c.mediator.Declare(message,c)
}

func (c *USA)GetMessage(message string)  {
	fmt.Println("USA得到消息",message)
}

type Iraq struct {
	mediator UnitedNations
}

func NewIraq(mediator UnitedNations) *Iraq {
	return &Iraq{mediator: mediator}
}

func (c *Iraq)Declare(message string)  {
	c.mediator.Declare(message,c)
}

func (c *Iraq)GetMessage(message string)  {
	fmt.Println("Iraq得到消息",message)
}

func main()  {
	UNSC := NewUnitedNationsSecurityCouncil()

	usa := NewUSA(UNSC)
	iraq := NewIraq(UNSC)

	UNSC.colleague1 = usa
	UNSC.colleague2 = iraq

	usa.Declare("不准研发核武器")
	iraq.Declare("我们不怕侵略")

}
