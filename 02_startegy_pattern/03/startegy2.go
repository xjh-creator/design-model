package main

import "fmt"

//运用简单工厂模式改写startegy1.go，
//缺点：简单工厂模式只是解决了对象的创建问题，但工厂本身包括了所有的收费方式，商城有经常更改打折额度和返利额度，每次更改优惠都会改动工厂，
//以致代码重新编译部署，这个方式有点糟糕。

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

type CashFactory struct {

}

func (c *CashFactory)createCashAccept(typeCash string,money float32) Cash {
	cashSuper := &CashSuper{money: money}
	switch typeCash {
	case "正常收费":
		return &CashNormal{cashSuper}
	case "满300返100":
		return &CashReturn{cashSuper,300,100}
	case "打8折":
		return &CashRebate{cashSuper,0.8}
	default:
		return nil
	}
}
var totalPrice float32

func main()  {
	factory := new(CashFactory)
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
		calculate := factory.createCashAccept(discount,A * float32(B))
		totalPrice += calculate.acceptCash()
		fmt.Println("总计金额为：",totalPrice)

		var operate string
		fmt.Println("请输入接下来操作(1重置、0退出程序、其它 继续计算)：")
		fmt.Scanf("%s",&operate)
		if operate == "1"{
			totalPrice = 0
		}else if operate == "0"{
			break
		}
	}

}





