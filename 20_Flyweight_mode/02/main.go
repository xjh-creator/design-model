package main

import "fmt"

type WebSite interface {
	Use()
}

type ConcreteWebSite struct {
	Name string
}

func NewConcreteWebSite(name string) *ConcreteWebSite {
	return &ConcreteWebSite{Name: name}
}

func (c *ConcreteWebSite)Use()  {
	fmt.Println("网站分类：",c.Name)
}

type WebSiteFactory struct{
	FlyWeights map[string]WebSite
}

func NewWebSiteFactory() *WebSiteFactory {
	return &WebSiteFactory{
		FlyWeights: make(map[string]WebSite),
	}
}

func (w *WebSiteFactory)GetWebSiteCategory(key string)WebSite {
	if w.FlyWeights[key] == nil{
		w.FlyWeights[key] = NewConcreteWebSite(key)
	}
	return w.FlyWeights[key]
}

func (w *WebSiteFactory)GetWebSiteCount()int {
	return len(w.FlyWeights)
}

func main()  {
	f := NewWebSiteFactory()

	fx := f.GetWebSiteCategory("产品展示")
	fx.Use()

	fy := f.GetWebSiteCategory("产品展示")
	fy.Use()

	fz := f.GetWebSiteCategory("产品展示")
	fz.Use()

	f1 := f.GetWebSiteCategory("博客")
	f1.Use()

	fm := f.GetWebSiteCategory("博客")
	fm.Use()

	fn := f.GetWebSiteCategory("博客")
	fn.Use()

	fmt.Println("网站分类总数为",f.GetWebSiteCount())
}

