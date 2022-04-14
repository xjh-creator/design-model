package main

import "fmt"

type Work struct {
	hour int
	finish bool
}

func (w *Work)WriteProgram()  {
	if w.hour < 12{
		fmt.Printf("当前时间：%d点 上午工作，精神百倍",w.hour)
	}else if w.hour < 13{
		fmt.Printf("当前时间：%d点 饿了，午饭；犯困，午休",w.hour)
	}else if w.hour < 17{
		fmt.Printf("当前时间：%d点 下午状态还不错，继续努力",w.hour)
	}else{
		if w.finish{
			fmt.Printf("当前时间：%d点 下班回家了",w.hour)
		}else{
			fmt.Printf("当前时间：%d点 不行了，睡着了",w.hour)
		}
	}
	fmt.Println()
}

func main()  {
	emergencyProjects := &Work{}
	emergencyProjects.hour = 9
	emergencyProjects.WriteProgram()
	emergencyProjects.hour = 10
	emergencyProjects.WriteProgram()
	emergencyProjects.hour = 12
	emergencyProjects.WriteProgram()
	emergencyProjects.hour = 13
	emergencyProjects.WriteProgram()
	emergencyProjects.hour = 14
	emergencyProjects.WriteProgram()
	emergencyProjects.hour = 17
	emergencyProjects.WriteProgram()

	emergencyProjects.finish = false

	emergencyProjects.WriteProgram()
	emergencyProjects.hour = 19
	emergencyProjects.WriteProgram()
	emergencyProjects.hour = 22
	emergencyProjects.WriteProgram()

}
