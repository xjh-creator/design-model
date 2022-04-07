package main

import "fmt"


type LeiFeng interface {
	Sweep()
	Wash()
	BuyRice()
}

type Person struct {
	name string
}

func (s *Person)Sweep()  {
	fmt.Println(s.name,"扫地")
}

func (s *Person)Wash()  {
	fmt.Println(s.name,"洗衣服")
}

func (s *Person)BuyRice()  {
	fmt.Println(s.name,"买米")
}

type Undergraduate struct {
	*Person
}

func NewUndergraduate() *Undergraduate{
	return &Undergraduate{&Person{"毕业生"}}
}

type Volunteer struct {
	*Person
}

func NewVolunteer() *Volunteer{
	return &Volunteer{&Person{"志愿者"}}
}

type Worker struct {
	*Person
}

func NewWorker() *Worker{
	return &Worker{&Person{"工人"}}
}

type LeiFengFactory interface {
	Create() LeiFeng
}

type UndergraduateFactory struct {

}

func (u *UndergraduateFactory)Create() LeiFeng {
	return NewUndergraduate()
}

type VolunteerFactory struct {

}

func (v *VolunteerFactory)Create() LeiFeng {
	return NewVolunteer()
}

type WorkerFactory struct {

}

func (w *WorkerFactory)Create() LeiFeng {
	return NewWorker()
}


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
