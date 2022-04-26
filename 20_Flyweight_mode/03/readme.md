###需求
为客户做网站，多个客户需求一致。享元模式正式版：每个客户的需求如果一致，即他们共享一个网站的实例

实体 | 说明
:---: | :---:
type WebSite interface| 网站抽象类，基类，也是享元类的基类
Use(user *User) | 方法，使用网站
type ConcreteWebSite struct | 具体网站，有一个字段Name，共享享元类
type User struct | 用户，非共享享元类
Use(user *User) | 方法，使用网站
type WebSiteFactory struct | 网站工厂，用来创建并管理享元对象
GetWebSiteCategory(key string) | 方法，得到网站，如果网站实例不存在即新建
GetWebSiteCount() | 方法，得到实例的个数



###测试：
    func main()  {
        f := NewWebSiteFactory()
    
        fx := f.GetWebSiteCategory("产品展示")
        fx.Use(NewUser("小菜"))
    
        fy := f.GetWebSiteCategory("产品展示")
        fy.Use(NewUser("小明"))
    
        fz := f.GetWebSiteCategory("产品展示")
        fz.Use(NewUser("小东"))
    
        f1 := f.GetWebSiteCategory("博客")
        f1.Use(NewUser("小红"))
    
        fm := f.GetWebSiteCategory("博客")
        fm.Use(NewUser("小祝"))
    
        fn := f.GetWebSiteCategory("博客")
        fn.Use(NewUser("小三"))
    
        fmt.Println("网站分类总数为",f.GetWebSiteCount())
    }

###享元模式定义
享元模式(Flyweight)，运用共享技术有效地支持大量细粒度对象。系统只使用少量对象，而这些对象都很相似，状态变化很小，可以实现对象地多次复用。由于享元模式要求能够共享地对象必须是细粒度对象，因此又称为轻量级模式，是一种结构型模式

###享元模式角色
1.Flyweight(抽象享元类)：一个接口或抽象类，声明了具体享元类的公共方法。
2.ConcreteFlyweight(具体享元类): 实现了抽象享元类，其实例称为享元对象。
3.UnsharedConcreteFlyweight(非共享具体享元类): 并不是所有的抽象享元类的子类都需要被共享，不能被共享的子类可设计为非共享具体享元类。
4.FlyweightFactory(享元工厂类): 用于创建并管理享元对象，一般设计为一个Key-Value键值对的集合(可以结合工厂模式设计)。其作用就在于：提供一个用于存储享元对象的享元池，当用户需要对象时，首先从享元池中获取，如果享元池中不存在，那么就创建一个新的享元对象返回给用户，并在享元池中保存该新增对象。

###应用场景
1. 一个应用程序使用了大量的对象，且对象造成了很大的存储开销
2. 对象又很多状态是外部状态，如果删除对象的外部状态，可以用较少的共享对象取代很多组对象

