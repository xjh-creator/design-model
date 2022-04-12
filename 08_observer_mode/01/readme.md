###需求
前台通知员工

实体 | 说明
:---: | :---:
type Secretary struct | 前台秘书（通知者）
Attach() | 添加需要通知的成员
Notify() | 通知需要通知的成员
SecretaryAction() | 设置动作
type StockObserver struct | 员工（被通知者）
NewStockObserver() | 创建一个新的被通知者
Update() | 被前台通知后采取行动

###测试：
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

}

###分析：
1. 这种写法耦合性太高，前台类与员工类互相耦合
2. 违背了开发-封闭原则
3. 违背了依赖倒转原则：让程序都依赖抽象，而不是相互依赖。

