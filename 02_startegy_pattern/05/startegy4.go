package main

import "fmt"
//运用了策略模式与简单工厂模式的结合
//简单工厂模式中，客户端需要认识两个对象，Cash 和 CashFactory
//策略模式与简单工厂模式的结合，只需要认识一个对象CashContext，耦合性更加降低

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
	CbxType string
}

//运用工厂方式来创建对象
func (c *CashContext)NewCash(total float32) Cash {
	var cash Cash
	switch c.CbxType {
	case "正常收费":
		cash =  &CashNormal{&CashSuper{money: total}}
	case "满300返100":
		cash = &CashReturn{&CashSuper{money: total},300,200}
	case "打8折":
		cash = &CashRebate{&CashSuper{money: total},0.8}
	default:
		cash = nil
	}
	return cash
}

func (c *CashContext)GetResult(total float32) float32 {
	return c.NewCash(total).acceptCash()
}

var total float32

func main()  {
	cashContext := new(CashContext)

	for{
		var A float32
		var B int
		fmt.Println("请输入单价：")
		fmt.Scanf("%f",&A)
		fmt.Println("请输入数量：")
		fmt.Scanf("%d",&B)
		fmt.Println("请输入打折方式(正常收费 满300减100 打八折)：")
		fmt.Scanf("%s",cashContext.CbxType)

		total += cashContext.GetResult(total)
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
