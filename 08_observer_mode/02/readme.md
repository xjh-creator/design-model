###需求
前台通知员工

实体 | 说明
:---: | :---:
type Secretary struct | 前台秘书（通知者）
Attach() | 添加需要通知的成员
Notify() | 通知需要通知的成员
SecretaryAction() | 设置动作
type Observer interface | 抽象观察者(基类)，用以跟前台类解耦
type Base struct | 把共同属性抽离出来
type StockObserver struct | 看股票的员工
NewStockObserver() | 创建一个新的看股票的员工
Update() | 被前台通知后采取行动
type NBAObserver struct | 看股票的员工
NewNBAObserver() | 创建一个新的NBA的员工
Update() | 被前台通知后采取行动

###测试：
    func main()  {
        //前台
        tongzizhe := &Secretary{}
        //看股票的同事
        tongshi1 := NewStockObserver("小明",tongzizhe)
        tongshi2 := NewNBAObserver("小东",tongzizhe)
    
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
1. 这种写法中前台秘书类与观察者类的耦合算是解除了，但前台秘书类算是通知者类的其中一种
2. 具体的观察者不应该去依赖具体的通知者，而是一个抽象的通知者
3. 通知者不只有一个，也可以有多个
