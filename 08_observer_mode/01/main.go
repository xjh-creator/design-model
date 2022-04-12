package main

import "fmt"

//Secretary前台秘书（通知者）
type Secretary struct {
	observers []*StockObserver
	action string
}
//Attach添加需要通知的成员
func (s *Secretary)Attach(observer *StockObserver)  {
	s.observers = append(s.observers,observer)
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
//StockObserver员工（被通知者）
type StockObserver struct {
	sub *Secretary
	name string
}
//NewStockObserver创建一个新的被通知者
func NewStockObserver(name string,sub *Secretary) *StockObserver {
	return &StockObserver{name: name,sub: sub}
}
//Update被前台通知后采取行动
func (s *StockObserver)Update()  {
	fmt.Printf("%s %s 关闭股票行情，继续工作",s.sub.SecretaryAction(""),s.name)
	fmt.Println()
}

func main()  {
	//前台
	tongzizhe := &Secretary{}
	//看股票的同事
	tongshi1 := NewStockObserver("小明",tongzizhe)
	tongshi2 := NewStockObserver("小东",tongzizhe)

	//前台记住两位同事
	tongzizhe.Attach(tongshi1)
	tongzizhe.Attach(tongshi2)
	//前台发现老板回来
	tongzizhe.SecretaryAction("老板回来了！")
	//通知两位同事
	tongzizhe.Notify()
}
