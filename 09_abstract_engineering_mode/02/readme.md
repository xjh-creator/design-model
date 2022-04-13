###需求
操作不同的数据库

实体 | 说明
:---: | :---:
type IUser interface | 基类，不同数据库操作的同一个类型的表
type User struct | 实体类
type SqlserverUser struct | Sql server用于操作User类
NewSqlserverUser() | 方法，初始化SqlserverUser
Insert(user *User) | 给数据库添加数据
GetUser(id int) | 得到数据库
type AccessUser struct | Access用于操作User类
NewSqlserverUser() | 方法，初始化AccessUser
Insert(user *User) | 给数据库添加数据
GetUser(id int) | 得到数据库
IFactory | 定义一个用于创建对象的接口，让子类决定实例化那一个类，工厂方法使一个类的实例化延迟到其子类
SqlServerFactory | 工厂SqlServer子类
CreateUser() | SqlServerFactory来实例化一个SqlserverUser
AccessFactory | 工厂Access子类
CreateUser() | AccessFactory来实例化一个AccessUser

###测试：
    func main()  {
       user := &User{ID: 1,Name: "xiaoming"}

        sqlServerFactory := &SqlServerFactory{}
        sqlServerUser := sqlServerFactory.CreateUser()
        sqlServerUser.Insert(user)
        sqlServerUser.GetUser(1)
    
        accessFactory := &AccessFactory{}
        accessUser := accessFactory.CreateUser()
        accessUser.Insert(user)
        accessUser.GetUser(1)
    }

###分析：
1. 用工厂方法模式改写模仿最基本的数据库访问程序
2. 更换数据库方便，只需要从声明并初始化SqlServerFactory改为AccessFactory即可

