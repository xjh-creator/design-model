###需求
操作不同的数据库

实体 | 说明
:---: | :---:
type User struct | 实体类
type SqlserverUser struct | 用于操作User类
Insert(user *User) | 给数据库添加数据
GetUser(id int) | 得到数据库

###测试：
    func main()  {
        user := &User{ID: 1,Name: "xiaoming"}
        su := NewSqlserverUser()
        su.Insert(user)
        su.GetUser(1)
    }

###分析：
1. 模仿最基本的数据库访问程序
2. 局限性大，不能更换数据库，SqlserverUser只能操作Sql server 上的User表

