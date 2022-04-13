package main

import "fmt"

//User表相关
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

//Class表相关
type IClass interface {
	Insert(class *Class)
	GetClass(id int)
}

//Class班级
type Class struct {
	ID int
	Name string
}

//SqlserverClass数据库 Sql server 的表Class
type SqlserverClass struct {
	classes []*Class
}

func NewSqlserverClass() *SqlserverClass {
	return &SqlserverClass{classes: make([]*Class,0)}
}

func (s *SqlserverClass)Insert(class *Class)  {
	fmt.Println("在SQL Server中给Class表增加一条纪录")
	s.classes = append(s.classes,class)
}

func (s *SqlserverClass)GetClass(id int)  {
	var result string
	for _,v := range s.classes{
		if v.ID == id{
			result = v.Name
		}
	}
	fmt.Println("在SQL Server中根据ID得到Class表一条纪录",result)
}

//AccessClass数据库 Access 的表Class
type AccessClass struct {
	classes []*Class
}

func NewAccessClass() *AccessClass {
	return &AccessClass{classes: make([]*Class,0)}
}

func (s *AccessClass)Insert(class *Class)  {
	fmt.Println("在Access中给Class表增加一条纪录")
	s.classes = append(s.classes,class)
}

func (s *AccessClass)GetClass(id int)  {
	var result string
	for _,v := range s.classes{
		if v.ID == id{
			result = v.Name
		}
	}
	fmt.Println("在Access中根据ID得到Class表一条纪录",result)
}

type DataAccess struct {

}

func (d *DataAccess)CreateUser(operate string) IUser {
	switch operate {
	case "Sqlserver":
		return NewSqlserverUser()
	case "Access":
		return NewAccessUser()
	default:
		return nil
	}
}

func (d *DataAccess)CreateClass(operate string) IClass {
	switch operate {
	case "Sqlserver":
		return NewSqlserverClass()
	case "Access":
		return NewAccessClass()
	default:
		return nil
	}
}
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
