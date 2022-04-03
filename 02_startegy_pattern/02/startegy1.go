package main

import "fmt"
//在startegy.go的基础上添加一个打折的功能
//如果要满足两种以上打折方式，这种改动方式就有点繁琐

//总计金额
var total1 float32

func makeDiscount(choice string,totalPrice float32) float32 {
	switch choice {
	case "1":
		return totalPrice * 0.8
	case "2":
		return totalPrice * 0.7
	case "3":
		return totalPrice * 0.5
	default:
		return totalPrice
	}

}

//商城简单收银软件
func main()  {
	var totalPrice float32
	for{
		var A float32
		var B int
		var discount string
		fmt.Println("请输入单价：")
		fmt.Scanf("%f",&A)
		fmt.Println("请输入数量：")
		fmt.Scanf("%d",&B)
		fmt.Println("请输入打折方式(1八折 2七折 3五折 其它正常收费)：")
		fmt.Scanf("%s",&discount)
		totalPrice = makeDiscount(discount,A * float32(B))
		total1 += totalPrice
		fmt.Println("总计金额为：",total1)

		var operate string
		fmt.Println("请输入接下来操作(1重置、0退出程序、其它 继续计算)：")
		fmt.Scanf("%s",&operate)
		if operate == "1"{
			total1 = 0
		}else if operate == "0"{
			break
		}
	}
}
