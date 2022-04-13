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
type DataAccess struct | 简单工厂类
CreateUser(operate string) | 创建一个IUser
CreateClass(operate string) | 创建一个IClass

###测试：
    func main()  {
        user := &User{ID: 1,Name: "xiaoming"}
        class := &Class{ID: 1,Name: "计科3班"}
    
        dataAccess := &DataAccess{}
        
        iUser := dataAccess.CreateUser("Sqlserver")
        iUser.Insert(user)
        iUser.GetUser(1)
        iClass := dataAccess.CreateClass("Sqlserver")
        iClass.Insert(class)
        iClass.GetClass(1)
    
        iUser = dataAccess.CreateUser("Access")
        iUser.Insert(user)
        iUser.GetUser(1)
        iClass = dataAccess.CreateClass("Access")
        iClass.Insert(class)
        iClass.GetClass(1)
    }

###分析：
1. 用简单工厂来改写抽象工厂

