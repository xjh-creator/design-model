package main

import "fmt"

type Graphics struct {

}

func (g *Graphics)DrawEllipse(p string)  {
	fmt.Printf("画笔的颜色是%s,画了头",p)
}

func (g *Graphics)DrawRectangle(p string)  {
	fmt.Printf("画笔的颜色是%s,画了身体",p)
}

func (g *Graphics)DrawLine(p,body string)  {
	fmt.Printf("画笔的颜色是%s,画了%s",p,body)
}


type ThinPerson struct {
	g Graphics
	Pen string
}

func (b *ThinPerson)Build()  {
	b.g.DrawRectangle(b.Pen)
	b.g.DrawRectangle(b.Pen)
	b.g.DrawLine(b.Pen,"瘦子的左手")
	b.g.DrawLine(b.Pen,"瘦子的右手")
	b.g.DrawLine(b.Pen,"瘦子的左腿")
	b.g.DrawLine(b.Pen,"瘦子的右腿")
}

type FatPerson struct {
	g Graphics
	Pen string
}

func (b *FatPerson)Build()  {
	b.g.DrawRectangle(b.Pen)
	b.g.DrawRectangle(b.Pen)
	b.g.DrawLine(b.Pen,"胖子的左手")
	b.g.DrawLine(b.Pen,"胖子的右手")
	b.g.DrawLine(b.Pen,"胖子的左腿")
	b.g.DrawLine(b.Pen,"胖子的右腿")
}

func main()  {
	thinPerson := &ThinPerson{Pen: "blue"}
	thinPerson.Build()

	fatPerson := &FatPerson{Pen: "red"}
	fatPerson.Build()
}
