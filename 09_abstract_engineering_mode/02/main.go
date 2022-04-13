package main

import "fmt"

type IUser interface {
	Insert(user *User)
	GetUser(id int)
}

//User用户
type User struct {
	ID int
	Name string
}

//SqlserverUser数据库 Sql server 的表User
type SqlserverUser struct {
	users []*User
}

func NewSqlserverUser() *SqlserverUser {
	return &SqlserverUser{users: make([]*User,0)}
}

func (s *SqlserverUser)Insert(user *User)  {
	fmt.Println("在SQL Server中给User表增加一条纪录")
	s.users = append(s.users,user)
}

func (s *SqlserverUser)GetUser(id int)  {
	var result string
	for _,v := range s.users{
		if v.ID == id{
			result = v.Name
		}
	}
	fmt.Println("在SQL Server中根据ID得到User表一条纪录",result)
}

//AccessUser数据库 Access 的表User
type AccessUser struct {
	users []*User
}

func NewAccessUser() *AccessUser {
	return &AccessUser{users: make([]*User,0)}
}

func (s *AccessUser)Insert(user *User)  {
	fmt.Println("在Access中给User表增加一条纪录")
	s.users = append(s.users,user)
}

func (s *AccessUser)GetUser(id int)  {
	var result string
	for _,v := range s.users{
		if v.ID == id{
			result = v.Name
		}
	}
	fmt.Println("在Access中根据ID得到User表一条纪录",result)
}

type IFactory interface {
	CreateUser() IUser
}

type SqlServerFactory struct {

}

func (s *SqlServerFactory)CreateUser() IUser {
	return NewSqlserverUser()
}

type AccessFactory struct {

}

func (s *AccessFactory)CreateUser() IUser {
	return NewAccessUser()
}

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
