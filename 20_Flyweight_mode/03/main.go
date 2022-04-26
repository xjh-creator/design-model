package main

import "fmt"

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

type WebSite interface {
	Use(user *User)
}

type ConcreteWebSite struct {
	Name string
}

func NewConcreteWebSite(name string) *ConcreteWebSite {
	return &ConcreteWebSite{Name: name}
}

func (c *ConcreteWebSite)Use(user *User)  {
	fmt.Println("网站分类：",c.Name," 用户：",user.Name)
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
	fx.Use(NewUser("小菜"))

	fy := f.GetWebSiteCategory("产品展示")
	fy.Use(NewUser("小明"))

	fz := f.GetWebSiteCategory("产品展示")
	fz.Use(NewUser("小东"))

	f1 := f.GetWebSiteCategory("博客")
	f1.Use(NewUser("小红"))

	fm := f.GetWebSiteCategory("博客")
	fm.Use(NewUser("小祝"))

	fn := f.GetWebSiteCategory("博客")
	fn.Use(NewUser("小三"))

	fmt.Println("网站分类总数为",f.GetWebSiteCount())
}

