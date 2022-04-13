package main

import "fmt"

//User用户
type User struct {
	ID int
	Name string
}

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

func main()  {
	user := &User{ID: 1,Name: "xiaoming"}
	su := NewSqlserverUser()
	su.Insert(user)
	su.GetUser(1)
}
