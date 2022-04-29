package main

import "fmt"

type IPerson interface {
	GetConclusion()
}

type Person struct {
	Name   string
	Action string
}

type Man struct {
	*Person
}

func NewMan() *Man {
	return &Man{&Person{Name: "男人"}}
}

func (m *Man) GetConclusion() {
	switch m.Action {
	case "成功":
		fmt.Printf("%s %s时，背后多半有一个伟大的女人。", m.Name, m.Action)
	case "失败":
		fmt.Printf("%s %s时，闷头喝酒，谁也不用劝。", m.Name, m.Action)
	case "恋爱":
		fmt.Printf("%s %s时，凡是不懂也要装懂。", m.Name, m.Action)
	}
	fmt.Println()
}

type Woman struct {
	*Person
}

func NewWonman() *Woman {
	return &Woman{&Person{Name: "女人"}}
}

func (m *Woman) GetConclusion() {
	switch m.Action {
	case "成功":
		fmt.Printf("%s %s时，背后多半有一个不成功的男人。", m.Name, m.Action)
	case "失败":
		fmt.Printf("%s %s时，眼泪汪汪，谁也劝不了。", m.Name, m.Action)
	case "恋爱":
		fmt.Printf("%s %s时，遇事情不懂也要装懂。", m.Name, m.Action)
	}
	fmt.Println()
}

func main() {
	persons := make([]IPerson, 0)

	man1 := NewMan()
	man1.Action = "成功"
	persons = append(persons, man1)
	woman1 := NewWonman()
	woman1.Action = "成功"
	persons = append(persons, woman1)

	man2 := NewMan()
	man2.Action = "失败"
	persons = append(persons, man2)
	woman2 := NewWonman()
	woman2.Action = "失败"
	persons = append(persons, woman2)

	man3 := NewMan()
	man3.Action = "恋爱"
	persons = append(persons, man3)
	woman3 := NewWonman()
	woman3.Action = "恋爱"
	persons = append(persons, woman3)

	for _, v := range persons {
		v.GetConclusion()
	}
}
