package main

import "fmt"

//IPlayer为目标角色，是客户调用的接口
type IPlayer interface {
	Attack()
	Defense()
}

//Translator为适配器角色,实现了客户调用的接口和另一接口
type Translator struct {
	Iadaptee
}

func NewTranslator(iadaptee Iadaptee) Translator {
	return Translator{iadaptee}
}

func (f Translator)Attack()  {
	f.进攻()
}

func (f Translator)Defense()  {
	f.防守()
}

//Iadaptee被转换角色，是需要被适配器转换成客户调用的接口
type Iadaptee interface {
	进攻()
	防守()
}

type ForeignCenter struct {
	name string
}

func (f ForeignCenter)进攻()  {
	fmt.Printf("外籍中锋%s进攻",f.name)
	fmt.Println()
}

func (f ForeignCenter)防守()  {
	fmt.Printf("外籍中锋%s防守",f.name)
	fmt.Println()
}

func NewAdaptee(name string) Iadaptee {
	return ForeignCenter{name}
}

func main()  {
	yaoming := NewAdaptee("yaoming")
	translator := NewTranslator(yaoming)
	translator.Attack()
	translator.Defense()
}

type Forwards struct {
	name string
}

func (f Forwards)Attack()  {
	fmt.Printf("前锋%s进攻",f.name)
	fmt.Println()
}

func (f Forwards)Defense()  {
	fmt.Printf("前锋%s防守",f.name)
	fmt.Println()
}

type Center struct {
	name string
}

func (f Center)Attack()  {
	fmt.Printf("中锋%s进攻",f.name)
	fmt.Println()
}

func (f Center)Defense()  {
	fmt.Printf("中锋%s防守",f.name)
	fmt.Println()
}

type Guards struct {
	name string
}

func (f Guards)Attack()  {
	fmt.Printf("后卫%s进攻",f.name)
	fmt.Println()
}

func (f Guards)Defense()  {
	fmt.Printf("后卫%s防守",f.name)
	fmt.Println()
}




