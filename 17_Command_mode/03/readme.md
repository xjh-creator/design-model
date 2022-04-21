###需求
演示路边烤羊肉串的实现。此版本使用松耦合，也是命令模式

实体 | 说明
:---: | :---:
type Command interface | 命令抽象类
ExcuteCommand() | 方法，执行命令
BakeName() string | 方法，得到菜名

实体 | 说明
:---: | :---:
type BakeMuttonCommand struct | 具体命令，烤羊肉
ExcuteCommand() | 方法，执行命令接收者的动作
BakeName() string | 方法，得到菜名
type BakeChickenWingCommand struct | 具体命令，烤鸡翅
ExcuteCommand() | 方法，执行命令接收者的动作
BakeName() string | 方法，得到菜名

实体 | 说明
:---: | :---:
type Barbecuer struct | 命令接收者，烤肉员
BakeMutton() | 方法，烤羊肉串
BakeChickenWing() | 方法，烤鸡翅

实体 | 说明
:---: | :---:
type Waiter struct | 服务员，记录命令并通知执行
SetOrder(command Command) | 方法，记录命令
CancelOrder(command Command) | 方法，删除
Notify() | 方法，通知执行命令


###测试：
    func main()  {
        //开店前的准备
        boy := NewBarbecuer()
        bakeMuttonCommand1 := NewBakeMuttonCommand(&boy)
        bakeMuttonCommand2 := NewBakeMuttonCommand(&boy)
        bakeChickenWingCommand1 := NewBakeChickenWingCommand(&boy)
        girl := NewWaiter()
    
        //开门营业,顾客点菜
        girl.SetOrder(bakeMuttonCommand1)
        girl.SetOrder(bakeMuttonCommand2)
        girl.SetOrder(bakeChickenWingCommand1)
    
        //点菜完毕，通知厨房
        girl.Notify()
    }

###命令模式定义
命令模式(Command)，将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化；对请求排队或记录请求日志，以及支持可撤销的操作

###角色
1. Receiver(接收者): 接收者执行与请求相关的操作，它具体实现对请求的业务处理。
2. Invoker(调用者)：请求发送者，通过命令对象来执行请求。
3. Command(抽象命令类)：一个抽象类或接口，声明了执行请求的Execute()方法，通过这些方法可以调用请求接收者的相关操作
4. ConcreteCommand(具体命令类)：具体命令类是抽象命令类的子类，实现了抽象命令类中声明的方法。在实现Execute()方法时，将调用接收者对象的相关操作(Action)

###出发点
命令模式可以将发送者和接收者完全解耦，发送者与接收者之间没有直接引用关系，发送请求的对象只需要知道如何发送请求，而不必知道如何完成请求。

###优点
1. 较为容易的设计一个命令队列
2. 在需要的情况下，能够将命令记入日志
3. 允许接收请求一方是否要否定请求
4. 可以容易地实现对请求的撤销和重做
5. 新增加的命令类不影响其它类

