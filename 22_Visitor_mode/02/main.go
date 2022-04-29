package main

import "fmt"

type Action interface {
	GetConclusion(person IPerson)
}

type Success struct {
}

func NewSuccess() *Success {
	return &Success{}
}

func (s *Success) GetConclusion(person IPerson) {
	switch person.GetName() {
	case "男人":
		fmt.Printf("%s %s时，背后多半有一个伟大的女人。", person.GetName(), "成功")
	case "女人":
		fmt.Printf("%s %s时，背后多半有一个不成功的男人。", person.GetName(), "成功")
	}
}

type Failing struct {
}

func NewFailing() *Failing {
	return &Failing{}
}

func (s *Failing) GetConclusion(person IPerson) {
	switch person.GetName() {
	case "男人":
		fmt.Printf("%s %s时，闷头喝酒，谁也不用劝。", person.GetName(), "失败")
	case "女人":
		fmt.Printf("%s %s时，眼泪汪汪，谁也劝不了。", person.GetName(), "失败")
	}
}

type Amativeness struct {
}

func NewAmativeness() *Amativeness {
	return &Amativeness{}
}

func (s *Amativeness) GetConclusion(person IPerson) {
	switch person.GetName() {
	case "男人":
		fmt.Printf("%s %s时，凡是不懂也要装懂。", person.GetName(), "恋爱")
	case "女人":
		fmt.Printf("%s %s时，遇事情不懂也要装懂。", person.GetName(), "恋爱")
	}
}

type IPerson interface {
	Accept(vistor Action)
	GetName() string
}

type Person struct {
	Name string
}

type Man struct {
	*Person
}

func NewMan() *Man {
	return &Man{&Person{Name: "男人"}}
}

func (m *Man) Accept(vistor Action) {
	vistor.GetConclusion(m)
}

func (m *Man) GetName() string {
	return m.Name
}

type Woman struct {
	*Person
}

func NewWonman() *Woman {
	return &Woman{&Person{Name: "女人"}}
}

func (m *Woman) Accept(vistor Action) {
	vistor.GetConclusion(m)
}

func (m *Woman) GetName() string {
	return m.Name
}

type ObjectStructure struct {
	Elements map[IPerson]IPerson
}

func NewObjectStructure() *ObjectStructure {
	return &ObjectStructure{make(map[IPerson]IPerson)}
}

func (o *ObjectStructure) Attach(element IPerson) {
	o.Elements[element] = element
}

func (o *ObjectStructure) Detach(element IPerson) {
	delete(o.Elements, element)
}

func (o *ObjectStructure) Display(visitor Action) {
	for v := range o.Elements {
		v.Accept(visitor)
		fmt.Println()
	}
}

func main() {
	o := NewObjectStructure()
	o.Attach(NewMan())
	o.Attach(NewWonman())

	v1 := NewSuccess()
	o.Display(v1)

	v2 := NewFailing()
	o.Display(v2)

	v3 := NewAmativeness()
	o.Display(v3)

}
