###需求
员工申请加薪或者请假，领导批准。责任链版本

实体 | 说明
:---: | :---:
type Request struct | 请求类
type Manager interface | 管理者，基类
SetSuperior(superior Manager) | 方法，设置上级
RequestApplication(request *Request) | 方法，处理请求
type CommonManger struct | 经理，具体管理者类
SetSuperior(superior Manager) | 经理的方法，设置上级
RequestApplication(request *Request) | 经理的处理请求方法，用自己的权限处理请求，超过权限的给上级处理
type Majordomo struct | 总监，具体管理者类
SetSuperior(superior Manager) | 总监的方法，设置上级
RequestApplication(request *Request) | 总监的处理请求方法，用自己的权限处理请求，超过权限的给上级处理
type GeneralManager struct | 总经理，具体管理者类
SetSuperior(superior Manager) | 总经理的方法，设置上级
RequestApplication(request *Request) | 总经理的处理请求方法，用自己的权限处理请求，超过权限的给上级处理

###测试：
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

###责任链定义
职责链模式(Chain Of Responsibility)：使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。将这个对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它为止。

###分析：
1. 处理请求时，请求会沿着链传递直到有一个具体处理对象来处理
2. 请求者跟处理者不耦合，且具体处理者自己并不需要知道链的结构，只需要保持一个指向后继者的引用。
