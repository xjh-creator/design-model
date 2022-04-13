###需求
操作不同的数据库

实体（产品等级1） | 说明
:---: | :---:
type IUser interface | 基类（用户表），不同数据库操作的同一个类型的表
type User struct | 实体类
type SqlserverUser struct | Sql server用于操作User类
NewSqlserverUser() | 方法，初始化SqlserverUser
Insert(user *User) | 给数据库添加数据
GetUser(id int) | 得到数据库
type AccessUser struct | Access用于操作User类
NewSqlserverUser() | 方法，初始化AccessUser
Insert(user *User) | 给数据库添加数据
GetUser(id int) | 得到数据库

实体（产品等级2） | 说明
:---: | :---:
type IClass interface | 基类（用户表），不同数据库操作的同一个类型的表
type Class struct | 实体类
type SqlserverClass struct | Sql server用于操作Class类
NewSqlserverClass() | 方法，初始化SqlserverClass
Insert(class *Class) | 给数据库添加数据
GetClass(id int) | 得到数据库
type AccessClass struct | Access用于操作Class类
NewSqlserverClass() | 方法，初始化AccessClass
Insert(class *Class) | 给数据库添加数据
GetClass(id int) | 得到数据库

实体（抽象工厂） | 说明
:---: | :---:
IFactory | 定义一个用于创建对象的接口，让子类决定实例化那一个类，工厂方法使一个类的实例化延迟到其子类
SqlServerFactory | 工厂SqlServer子类
CreateUser() | SqlServerFactory来实例化一个SqlserverUser
CreateClass() | SqlServerFactory来实例化一个SqlserverClass
AccessFactory | 工厂Access子类
CreateUser() | AccessFactory来实例化一个AccessUser
CreateClass() | AccessFactory来实例化一个SqlserverClass

###测试：
    func main()  {
        var iFactory IFactory
        var iClass IClass
        var iUser IUser
    
        user := &User{ID: 1,Name: "xiaoming"}
        class := &Class{ID: 1,Name: "计科3班"}
    
        iFactory = &SqlServerFactory{}
        iUser = iFactory.CreateUser()
        iUser.Insert(user)
        iUser.GetUser(1)
        iClass = iFactory.CreateClass()
        iClass.Insert(class)
        iClass.GetClass(1)
    
        iFactory = &AccessFactory{}
        iUser = iFactory.CreateUser()
        iUser.Insert(user)
        iUser.GetUser(1)
        iClass = iFactory.CreateClass()
        iClass.Insert(class)
        iClass.GetClass(1)
    }

###分析：
1. 用抽象工厂方法模式来改写02用的工厂方法模式
2. 当系统所提供的工厂所需生产的具体产品并不是一个简单的对象，而是多个位于不同产品等级结构中属于不同类型的具体产品时需要使用抽象工厂模式。
3. 此系统中有两个数据库Access和sql server,每个数据库有各有两张表 user,class,这种属于不同产品等级结构的不同类型的具体产品
4. 抽象工厂方法模式与工厂方法模式的区别就是工厂方法模式只有一个产品等级结构，比如一种user表，数据库Access,sql server
5. 好处：横向切换十分便利，比如SqlServerFactory 切换到 AccessFactory
6. 坏处：纵向扩展十分麻烦，比如再增加一种表college,除了自身所需要的结构，还需要改动SqlServerFactory 和 AccessFactory

