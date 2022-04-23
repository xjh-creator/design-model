###需求
员工申请加薪或者请假，领导批准。责任链版本

实体 | 说明
:---: | :---:
type UnitedNations interface | 联合国组织，中介者基类
Declare(message string,country Country) | 方法，接收一个国家a的信息，然后向国家b发送国家a发送的信息
type UnitedNationsSecurityCouncil struct | 联合国安全理事会，中介者具体类
Declare(message string,country Country) | 方法，接收一个国家a的信息，然后向国家b发送国家a发送的信息

实体 | 说明
:---: | :---:
type Country interface | 国家，基类
Declare(message string) | 方法，国家a发送要发送的信息，中介者接收后发给国家b
GetMessage(message string) | 方法，得到信息
type USA struct | 美国，具体国家
Declare(message string) | 方法，美国发送要发送的信息，中介者接收后发给国家b
GetMessage(message string) | 方法，美国得到信息
type Iraq struct | 伊拉克，具体国家
Declare(message string) | 方法，伊拉克发送要发送的信息，中介者接收后发给国家b
GetMessage(message string) | 方法，伊拉克得到信息

###测试：
    func main()  {
        UNSC := NewUnitedNationsSecurityCouncil()

        usa := NewUSA(UNSC)
        iraq := NewIraq(UNSC)
    
        UNSC.colleague1 = usa
        UNSC.colleague2 = iraq
    
        usa.Declare("不准研发核武器")
        iraq.Declare("我们不怕侵略")
    }

###中介者模式定义
中介者模式(Mediator): 用一个中介对象来封装一系列的对象交互。中介者使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间地交互

###解决问题
对于一个模块，可能由很多对象构成，而且这些对象之间可能存在相互的引用，为了减少对象两两之间复杂的引用关系，使之称为一个松耦合的系统，这就是中介者模式的模式动机

###角色
1. Mediator(抽象中介者)：它定义了一个接口，该接口用于与各同事对象之间进行通信。
2. ConcreteMediator(具体中介者)：它实现了接口，通过协调各个同事对象来实现协作行为，维持各个同事对象的引用
3. Colleague(抽象同事类)：它定义了各个同事类公有的方法，并声明了一些抽象方法来供子类实现，同时维持了一个对抽象中介者类的引用，其子类可以通过该引用来与中介者通信。
4. ConcreteColleague(具体同事类)：抽象同事类的子类，每一个同事对象需要和其他对象通信时，都需要先与中介者对象通信，通过中介者来间接完成与其他同事类的通信

###分析：
1. 中介者的出现减少各个同事的耦合，可以独立改变和复用中介者和同事
2. 由于具体的中介者控制了集中化，交互的复杂性变为了中介者复杂性，使得中介者比任何一个同事都复杂
3. 中介者模式一般应用于一组对象以定义良好但是复杂的方式进行通信的场合，以及定制一个分布在各个类中的行为，但又不想生成太多子类的场合。
