package main

import "fmt"

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
	SecretaryAction(action string) string
}

//Secretary前台秘书（通知者）
type Secretary struct {
	observers []Observer
	action string
}
//Attach添加需要通知的成员
func (s *Secretary)Attach(observer Observer)  {
	s.observers = append(s.observers,observer)
}
//Detach删除需要通知的成员
func (s *Secretary)Detach(observer Observer)  {
	for i,v := range s.observers{
		if v == observer{
			s.observers = append(s.observers[:i],s.observers[i+1:]...)
			break
		}
	}
}
//Notify通知需要通知的成员
func (s *Secretary)Notify()  {
	for i,_ := range s.observers{
		s.observers[i].Update()
	}
}
//SecretaryAction设置动作
func (s *Secretary)SecretaryAction(action string) string {
	if action != ""{
		s.action = action
	}
	return s.action
}

//Secretary前台秘书（通知者）
type Boss struct {
	observers []Observer
	action string
}
//Detach删除需要通知的成员
func (s *Boss)Detach(observer Observer)  {
	for i,v := range s.observers{
		if v == observer{
			s.observers = append(s.observers[:i],s.observers[i+1:]...)
			break
		}
	}
}
//Attach添加需要通知的成员
func (s *Boss)Attach(observer Observer)  {
	s.observers = append(s.observers,observer)
}
//Notify通知需要通知的成员
func (s *Boss)Notify()  {
	for i,_ := range s.observers{
		s.observers[i].Update()
	}
}
//SecretaryAction设置动作
func (s *Boss)SecretaryAction(action string) string {
	if action != ""{
		s.action = action
	}
	return s.action
}

//Observer抽象观察者，用以跟前台类解耦
type Observer interface {
	Update()
}
//Base，把共同属性抽离出来
type Base struct {
	sub Subject
	name string
}
//StockObserver看股票的员工（被通知者）
type StockObserver struct {
	*Base
}
//NewStockObserver创建一个新的被通知者
func NewStockObserver(name string,sub Subject) *StockObserver {
	return &StockObserver{&Base{name: name,sub: sub}}
}
//Update被前台通知后采取行动
func (s *StockObserver)Update()  {
	fmt.Printf("%s %s 关闭股票行情，继续工作",s.sub.SecretaryAction(""),s.name)
	fmt.Println()
}
//NBAObserver看股票的员工（被通知者）
type NBAObserver struct {
	*Base
}
//NewStockObserver创建一个新的被通知者
func NewNBAObserver(name string,sub Subject) *NBAObserver {
	return &NBAObserver{&Base{name: name,sub: sub}}
}
//Update被前台通知后采取行动
func (s *NBAObserver)Update()  {
	fmt.Printf("%s %s 关闭NBA直播，继续工作",s.sub.SecretaryAction(""),s.name)
	fmt.Println()
}

func main()  {
	//前台
	tongzizhe := &Secretary{}
	boss := &Boss{}
	//看股票的同事
	tongshi1 := NewStockObserver("小明",tongzizhe)
	tongshi2 := NewNBAObserver("小东",boss)

	//前台记住两位同事
	tongzizhe.Attach(tongshi1)
	boss.Attach(tongshi2)
	//前台发现老板回来
	tongzizhe.SecretaryAction("老板回来了！")
	boss.SecretaryAction("我回来了")
	//通知两位同事
	tongzizhe.Notify()
	boss.Notify()
}
