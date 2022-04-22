package main

import "fmt"

//Request请求
type Request struct {
	RequestType string
	RequestContent string
	Number int
}

func NewRequest(requestType,content string,number int) *Request {
	return &Request{
		RequestType: requestType,
		RequestContent: content,
		Number: number,
	}
}

//Manager管理者，基类
type Manager interface {
	SetSuperior(superior Manager)
	RequestApplication(request *Request)
}

//CommonManger经理
type CommonManger struct {
	Superior Manager
	Name string
}

func NewCommonManger(name string) *CommonManger {
	return &CommonManger{Name: name}
}

func (c *CommonManger)SetSuperior(superior Manager)  {
	c.Superior = superior
}

func (c *CommonManger)RequestApplication(request *Request){
	if request.RequestType == "请假" && request.Number <= 2{
		fmt.Printf("%s : %s 数量 %d 被批准",c.Name,request.RequestContent,request.Number)
		fmt.Println()
	}else{
		if c.Superior != nil{
			c.Superior.RequestApplication(request)
		}
	}
}
//Majordomo总监
type Majordomo struct {
	Superior Manager
	Name string
}

func NewMajordomo(name string) *Majordomo {
	return &Majordomo{Name: name}
}

func (c *Majordomo)SetSuperior(superior Manager)  {
	c.Superior = superior
}

func (m *Majordomo)RequestApplication(request *Request){
	if request.RequestType == "请假" && request.Number <= 5{
		fmt.Printf("%s : %s 数量 %d 被批准",m.Name,request.RequestContent,request.Number)
		fmt.Println()
	}else{//如果存在上级，转给上级处理
		if m.Superior != nil{
			m.Superior.RequestApplication(request)
		}
	}
}

//GeneralManager总经理
type GeneralManager struct {
	Superior Manager
	Name string
}

func NewGeneralManager(name string) *GeneralManager {
	return &GeneralManager{Name: name}
}

func (c *GeneralManager)SetSuperior(superior Manager)  {
	c.Superior = superior
}

func (g *GeneralManager)RequestApplication(request *Request){
	if request.RequestType == "请假" && request.Number <= 5{//总经理准许下属任意天数的假期
		fmt.Printf("%s : %s 数量 %d 被批准",g.Name,request.RequestContent,request.Number)
		fmt.Println()
	}else if request.RequestType == "加薪" && request.Number <= 500{//加薪再500以内，没有问题
		fmt.Printf("%s : %s 数量 %d 我无权先处理",g.Name,request.RequestContent,request.Number)
		fmt.Println()
	}else if request.RequestType == "加薪" && request.Number > 500{//加薪超过500,就要考虑一下
		fmt.Printf("%s : %s 数量 %d 再说吧",g.Name,request.RequestContent,request.Number)
		fmt.Println()
	}
}

func main()  {
	//三个管理者
	jinli := NewCommonManger("金利")
	zongjian := NewMajordomo("宗剑")
	zhongjingli := NewGeneralManager("钟精励")
	//设置上级
	jinli.SetSuperior(zongjian)
	zongjian.SetSuperior(zhongjingli)
    //四个请求
	request1 := NewRequest("请假","小菜请求请假",1)
	request2 := NewRequest("请假","小菜请求请假",4)
	request3 := NewRequest("加薪","小菜请求加薪",500)
	request4 := NewRequest("加薪","小菜请求加薪",1000)
	//经理处理请求
	jinli.RequestApplication(request1)
	jinli.RequestApplication(request2)
	jinli.RequestApplication(request3)
	jinli.RequestApplication(request4)

}
