package main

import "fmt"

//ProductPerson 为产品
type ProductPerson struct {
	head string
	body string
	leftHand string
	rightHand string
	leftLeg string
	rightLeg string
}

func (p ProductPerson)show()  {
	fmt.Println("角色的头：",p.head)
	fmt.Println("角色的身体：",p.body)
	fmt.Println("角色的左手：",p.leftHand)
	fmt.Println("角色的右手：",p.rightHand)
	fmt.Println("角色的左腿：",p.leftLeg)
	fmt.Println("角色的右腿：",p.rightLeg)
}

//Builder 抽象建造者类
type Builder interface {
	BuilderHead()
	BuilderBody()
	BuilderLeftHand()
	BuilderRightHand()
	BuilderLeftLeg()
	BuilderRightLeg()
	CreateProductPerson() ProductPerson
}

//ThinPerson具体建造者类
type ThinPersonBuilder struct {
	thinPerson ProductPerson
}

func (t *ThinPersonBuilder)BuilderHead(){
	t.thinPerson.head = "瘦子的头"
}

func (t *ThinPersonBuilder)BuilderBody(){
	t.thinPerson.body = "瘦子的身体"
}

func (t *ThinPersonBuilder)BuilderLeftHand(){
	t.thinPerson.leftHand = "瘦子的左手"
}

func (t *ThinPersonBuilder)BuilderRightHand(){
	t.thinPerson.rightHand = "瘦子的右手"
}

func (t *ThinPersonBuilder)BuilderLeftLeg(){
	t.thinPerson.leftLeg = "瘦子的左腿"
}

func (t *ThinPersonBuilder)BuilderRightLeg(){
	t.thinPerson.rightLeg = "瘦子的右腿"
}

func (t ThinPersonBuilder)CreateProductPerson() ProductPerson {
	return t.thinPerson
}

//指挥者
type Director struct {
	builder Builder
}

//Construct指挥建造过程
func (d *Director)Construct() ProductPerson {
	d.builder.BuilderLeftHand()
	d.builder.BuilderHead()
	d.builder.BuilderBody()
	d.builder.BuilderRightHand()
	d.builder.BuilderLeftLeg()
	d.builder.BuilderRightLeg()
	return d.builder.CreateProductPerson()
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}


func main()  {
	thinPersonBuilder := ThinPersonBuilder{}

	director := NewDirector(&thinPersonBuilder)
	thinPerson := director.Construct()
	thinPerson.show()
}
