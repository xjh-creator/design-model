package main

import "fmt"

var (
	Hour = 0
	WorkFinished = false
)

func WriteProgram()  {
	if Hour < 12{
		fmt.Printf("当前时间：%d点 上午工作，精神百倍",Hour)
	}else if Hour < 13{
		fmt.Printf("当前时间：%d点 饿了，午饭；犯困，午休",Hour)
	}else if Hour < 17{
		fmt.Printf("当前时间：%d点 下午状态还不错，继续努力",Hour)
	}else{
		if WorkFinished{
			fmt.Printf("当前时间：%d点 下班回家了",Hour)
		}else{
			fmt.Printf("当前时间：%d点 不行了，睡着了",Hour)
		}
	}
	fmt.Println()
}

func main()  {
	Hour = 9
	WriteProgram()
	Hour = 10
	WriteProgram()
	Hour = 12
	WriteProgram()
	Hour = 13
	WriteProgram()
	Hour = 14
	WriteProgram()
	Hour = 17
	WriteProgram()

	WorkFinished = true

	WriteProgram()
	Hour = 19
	WriteProgram()
	Hour = 22
	WriteProgram()

}
