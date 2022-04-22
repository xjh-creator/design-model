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

type Manager struct {
	Name string
}

func NewManager(name string) *Manager {
	return &Manager{
		Name: name,
	}
}

func (m *Manager)GetResult(managerLevel string,request *Request)  {
	switch managerLevel {
	case "经理":
		if request.RequestType == "请假" && request.Number <= 2{
			fmt.Printf("%s : %s 数量 %d 被批准",m.Name,request.RequestContent,request.Number)
			fmt.Println()
		}else{
			fmt.Printf("%s : %s 数量 %d 我无权先处理",m.Name,request.RequestContent,request.Number)
			fmt.Println()
		}
	case "总监":
		if request.RequestType == "请假" && request.Number <= 5{
			fmt.Printf("%s : %s 数量 %d 被批准",m.Name,request.RequestContent,request.Number)
			fmt.Println()
		}else{
			fmt.Printf("%s : %s 数量 %d 我无权先处理",m.Name,request.RequestContent,request.Number)
			fmt.Println()
		}
	case "总经理":
		if request.RequestType == "请假" && request.Number <= 5{
			fmt.Printf("%s : %s 数量 %d 被批准",m.Name,request.RequestContent,request.Number)
			fmt.Println()
		}else if request.RequestType == "加薪" && request.Number <= 500{
			fmt.Printf("%s : %s 数量 %d 我无权先处理",m.Name,request.RequestContent,request.Number)
			fmt.Println()
		}else if request.RequestType == "加薪" && request.Number > 500{
			fmt.Printf("%s : %s 数量 %d 再说吧",m.Name,request.RequestContent,request.Number)
			fmt.Println()
		}

	}
}

func main()  {
	//三个管理者
	jinli := NewManager("金利")
	zongjian := NewManager("宗剑")
	zhongjingli := NewManager("钟精励")

	request1 := NewRequest("加薪","小菜请求加薪",1000)

	jinli.GetResult("经理",request1)
	zongjian.GetResult("总监",request1)
	zhongjingli.GetResult("总经理",request1)

	request2 := NewRequest("请假","小菜请求请假",3)

	jinli.GetResult("经理",request2)
	zongjian.GetResult("总监",request2)
	zhongjingli.GetResult("总经理",request2)

}
