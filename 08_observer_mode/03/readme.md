###需求
前台通知员工

实体 | 说明
:---: | :---:
type Subject interface | 抽象通知者（基类）
type Secretary struct | 前台秘书（通知者）
Attach() | 添加需要通知的成员
Notify() | 通知需要通知的成员
SecretaryAction() | 设置动作
type Boss struct | 老板（通知者）
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

}

###分析：
1. 这种写法就是观察者模式（又称发布-订阅模式）
2. 观察者的角色有（1）观察者类，（2）具体观察者，（3）主题类（通知者），（4）具体通知者
3. 主题类：每个主题有多个观察者，抽象主题可以定义一个接口来增加和删除观察者对象
4. 具体主题：将有关状态保存到具体观察者对象，当具体主题的内部状态改变时，通知所有登记过的观察者
5. 抽象观察者：为具体观察者定义一个接口，在得到主题通知时更新自己。
6. 具体观察者：实现抽象观察者角色所要求的更新接口，以便本身的状态与主题的状态相协调。
