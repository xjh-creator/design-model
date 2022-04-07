实体 | 说明
:---: | :---:
type LeiFeng interface | 雷锋（可理解基类），是Person、Undergraduate、Volunteer、Worker的源头
type Person struct | 人，有三个学雷锋的动作Sweep、Wash、BuyRice
type Undergraduate struct | 毕业生，继承Person与其方法
type Volunteer struct | 志愿者，继承Person与其方法
type Worker struct | 工人，继承Person与其方法
type LeiFengFactory interface | 定义一个用于创建对象的接口，让子类决定实例化那一个类，工厂方法使一个类的实例化延迟到其子类
type UndergraduateFactory struct | 工厂毕业生子类，用来实例化Undergraduate
type VolunteerFactory struct | 工厂志愿者子类，用来实例化Volunteer
type WorkerFactory struct | 工厂工人子类，用来实例化Worker


###测试：
    func main()  {
        var leiFengFactory LeiFengFactory
        var leiFeng LeiFeng
    
        leiFengFactory = &UndergraduateFactory{}
        leiFeng = leiFengFactory.Create()
        leiFeng.Wash()
        leiFeng.Sweep()
        leiFeng.BuyRice()
    
        leiFengFactory = &VolunteerFactory{}
        leiFeng = leiFengFactory.Create()
        leiFeng.Wash()
        leiFeng.Sweep()
        leiFeng.BuyRice()
    
        leiFengFactory = &WorkerFactory{}
        leiFeng = leiFengFactory.Create()
        leiFeng.Wash()
        leiFeng.Sweep()
        leiFeng.BuyRice()
    }


###分析
1. 在工厂方法模式中，工厂用来创建客户端所需要的对象，而创建细节都隐藏起来，达到比较好的封装效果
2. 基于工厂角色和产品角色的多态性设计是工厂方法模式的关键。使工厂自主创建产品对象。之所以有多态工厂模式的称呼，就是所有具体工厂有一个统一父类。
3. 在系统扩展功能时，不需要修改原有类的方法，而是通过增加代码的方式来扩展功能，比较好的遵循了开放-封闭原则。
4. 工厂方法比简单工厂方法的好处就是克服简单工厂违背开放-封闭原则的缺点，有保持了封装对象创建过程的优点。
5. 不好的地方就是增加一个产品就需要增加一个产品工厂的类，增加了额外的开发量。