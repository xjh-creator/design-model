package main

import "fmt"
//运用了策略模式，当这里又存在了一个问题，就是需要客户端来判断用哪个算法
//所以更好的解决方式是简单工厂模式跟策略模式相结合

type Cash interface {
	acceptCash() float32
}

//收费
type CashSuper struct {
	money float32
}

//正常收费
type CashNormal struct {
	*CashSuper
}

func (c *CashNormal)acceptCash()float32{
	return c.money
}

//打折收费
type CashRebate struct {
	*CashSuper
	MoneyRebate float32
}

func (c *CashRebate)acceptCash()float32 {
	return c.money * c.MoneyRebate
}

//返利收费
type CashReturn struct {
	*CashSuper
	MoneyCondition,MoneyReturn float32
}

//moneyCondition满足金额的条件,moneyReturn返回的金额
func (c *CashReturn)acceptCash() float32{
	if c.money > c.MoneyCondition{
		c.money -= c.MoneyReturn
	}
	return c.money
}

type CashContext struct {
	cs Cash
}


func (c *CashContext)GetResult() float32 {
	return c.cs.acceptCash()
}

var total float32

func main()  {
	cashContext := new(CashContext)

	for{
		var A float32
		var B int
		var discount string
		fmt.Println("请输入单价：")
		fmt.Scanf("%f",&A)
		fmt.Println("请输入数量：")
		fmt.Scanf("%d",&B)
		fmt.Println("请输入打折方式(正常收费 满300减100 打八折)：")
		fmt.Scanf("%s",&discount)

		switch discount {
		case "正常收费":
			cashContext.cs = &CashNormal{&CashSuper{money: A*float32(B)}}
		case "满300返100":
			cashContext.cs = &CashReturn{&CashSuper{money: A*float32(B)},300,200}
		case "打8折":
			cashContext.cs = &CashRebate{&CashSuper{money: A*float32(B)},0.8}
		default:
			cashContext.cs = nil
		}

		total += cashContext.GetResult()
		fmt.Println("总计金额为：",total)

		var operate string
		fmt.Println("请输入接下来操作(1重置、0退出程序、其它 继续计算)：")
		fmt.Scanf("%s",&operate)
		if operate == "1"{
			total = 0
		}else if operate == "0"{
			break
		}
	}
}