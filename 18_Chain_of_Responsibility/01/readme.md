###需求
员工申请加薪或者请假，领导批准。简单版本

实体 | 说明
:---: | :---:
type Request struct | 请求类
type Manager struct | 管理者类
GetResult(managerLevel string,request *Request) | 方法，得到请求结果

###测试：
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

###分析：
1. GetResult()方法里面结果比较长，且太多分支判断。违背了单一职职责，开放-封闭原则
2. 重构思路：把管理者的类别做成管理者的子类，利用多态性来化解分支带来的僵化。再利用责任链模式来解决关联（如：经理无权上报总监，总监无权上报总经理）

