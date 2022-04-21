###需求
演示路边烤羊肉串的实现。此版本使用松耦合设计

实体 | 说明
:---: | :---:
type Command interface | 命令抽象类
ExcuteCommand() | 方法，执行命令

实体 | 说明
:---: | :---:
type BakeMuttonCommand struct | 具体命令，烤羊肉
ExcuteCommand() | 方法，执行命令接收者的动作
type BakeChickenWingCommand struct | 具体命令，烤鸡翅
ExcuteCommand() | 方法，执行命令接收者的动作

实体 | 说明
:---: | :---:
type Barbecuer struct | 命令接收者，烤肉员
BakeMutton() | 方法，烤羊肉串
BakeChickenWing() | 方法，烤鸡翅

实体 | 说明
:---: | :---:
type Waiter struct | 服务员，纪录命令并通知执行
SetOrder(command Command) | 方法，纪录命令
Notify() | 方法，通知执行命令


###测试：
    func main()  {
        //开店前的准备
        boy := NewBarbecuer()
        bakeMuttonCommand1 := NewBakeMuttonCommand(&boy)
        bakeMuttonCommand2 := NewBakeMuttonCommand(&boy)
        bakeChickenWingCommand1 := NewBakeChickenWingCommand(&boy)
        girl := NewWaiter()
    
        //开门营业
        girl.SetOrder(bakeMuttonCommand1)
        girl.Notify()
        girl.SetOrder(bakeMuttonCommand2)
        girl.Notify()
        girl.SetOrder(bakeChickenWingCommand1)
        girl.Notify()
    }

###分析
1. 这种写法比紧耦合设计改进许多，但仍存在很多问题
2. 服务员不单一次值纪录客户点的一个菜，就去通知，而是纪录点完多个菜才去通知制作
3. 如果食材没有了，应该通知客户
4. 客户点的菜需要纪录，以备收费，以及后期统计
5. 客户有时可能会取消一些菜

