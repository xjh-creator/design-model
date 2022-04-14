package main

import "fmt"


type Work struct {
	hour int
	finish bool
	current State
}

func (w *Work)SetHour(hour int)  {
	w.hour = hour
}

func (w *Work)SetState(s State)  {
	w.current = s
}

func (w *Work)SetFinish(finish bool)  {
	w.finish = finish
}

func (w Work)WriteProgram()  {
	w.current.WriteProgram(w)
}

type State interface {
	WriteProgram(w Work)
}

//ForenoonState上午工作状态
type ForenoonState struct {

}

func (f *ForenoonState)WriteProgram(w Work) {
	if w.hour < 12{
		fmt.Printf("当前时间：%d点 上午工作，精神百倍",w.hour)
	}else{
		//超过十二点，转到中午工作模式
		w.SetState(new(NoonState))
		w.WriteProgram()
	}
	fmt.Println()
}
//NoonState中午工作状态
type NoonState struct {

}

func (f *NoonState)WriteProgram(w Work) {
	if w.hour < 13{
		fmt.Printf("当前时间：%d点 饿了，午饭；犯困，午休",w.hour)
	}else{
		//超过十三点，转到下午工作模式
		w.SetState(&AfterNoonState{})
		w.WriteProgram()
	}
	fmt.Println()
}
//AfterNoonState下午工作状态
type AfterNoonState struct {

}

func (f *AfterNoonState)WriteProgram(w Work) {
	if w.hour < 17{
		fmt.Printf("当前时间：%d点 下午状态还不错，继续努力",w.hour)
	}else{
		//超过十七点，转到傍晚工作模式
		w.SetState(&EveningState{})
		w.WriteProgram()
	}
	fmt.Println()
}
//EveningState晚上工作状态
type EveningState struct {

}

func (f *EveningState)WriteProgram(w Work) {
	if w.finish{
		w.SetState(&RestState{})
		w.WriteProgram()
	}else{
		if w.hour < 21{
			fmt.Printf("当前时间：%d点 加班哦，疲惫至极",w.hour)
		}else{
			w.SetState(&SleepingState{})
			w.WriteProgram()
		}
	}
	fmt.Println()
}
//SleepingState睡眠状态
type SleepingState struct {

}

func (f *SleepingState)WriteProgram(w Work) {
	fmt.Printf("当前时间：%d点 不行了，睡着了",w.hour)
	fmt.Println()
}
//RestState休息状态
type RestState struct {

}

func (f *RestState)WriteProgram(w Work) {
	fmt.Printf("当前时间：%d点 下班回家了",w.hour)
	fmt.Println()
}



func main()  {
	emergencyProjects := new(Work)
	tState := &NoonState{}
	emergencyProjects.SetState(tState)
	emergencyProjects.SetFinish(true)
	emergencyProjects.SetHour(15)
	emergencyProjects.WriteProgram()
}
