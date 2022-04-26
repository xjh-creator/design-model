package main

import "fmt"

type WebSite struct {
	Name string
}

func NewWebSite(name string) *WebSite {
	return &WebSite{
		Name: name,
	}
}

func (u WebSite)Use()  {
	fmt.Println("网站分类：",u.Name)
}

func main()  {
	fx := NewWebSite("产品展示")
	fx.Use()
	fy := NewWebSite("产品展示")
	fy.Use()
	fz := NewWebSite("产品展示")
	fz.Use()

	fl := NewWebSite("博客")
	fl.Use()
	fm := NewWebSite("博客")
	fm.Use()
	fn := NewWebSite("博客")
	fn.Use()
}
