#策略模式
策略模式是一种定义一系列算法的方法，所有算法完成的是任务，只是实现不同方式不同，它可以以相同的方式调用所有的算法。
减少各种算法类与使用算法类直接的耦合。

策略模式：
1. 定义了一组算法（业务规则）；
2. 封装了每个算法；
3. 这族的算法可互换代替（interchangeable）。

优点：
1. strategy类层次为Context定义一系列的可重用的算法或者行为
2. 继承有助于析取出这些算法中的公共功能
3. 比如加减乘除四个策略中，共同的功能就是计算的结果GetResult,所以有了抽象的类计算方式
4. 策略模式就是来封装算法的

#不运用简单工程模式的方式
    var A,B,result int
	var operate string
	fmt.Println("请输入数字A：")
	fmt.Scanf("%d",&A)
	fmt.Println("请选择运算符（+、-、*、/）：")
	fmt.Scanf("%s",&operate)
	fmt.Println("请输入数字B：")
	fmt.Scanf("%d",&B)
	switch operate {
	case "+":
		result = A + B
	case "-":
		result = A - B
	case "*":
		result = A * B
	case "/":
		if B == 0{
			result = -1
			break
		}
		result = A / B
	default:
		result = -1
	}
	fmt.Println("计算结果为（-1表示有错）：",result)

这里可以看出一个很严重的问题就是：代码耦合度高，不容易维护、不容易扩展、不容易复用

#运用策略模式
    /**
    * 策略接口
     */
    type Cash interface {
    acceptCash() float32
    }

    //收费
    type CashSuper struct {
    money float32
    }

    /**
    * CashNormal 极其方法acceptCash就是我们定义的一个策略
    */
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

    /**
    * CashContext 用来维护一个对Strategy对象的引用
    * CashContext 传入一个具体的策略对象，然后根据策略对象，调用其算法的方法
    */
    type CashContext struct {
        cs Cash
    }

    func (c *CashContext)GetResult() float32 {
        return c.cs.acceptCash()
    }
测试：
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
在测试中，可以看到需要客户端来判断使用哪个策略，更好的解决方案是策略模式跟工厂模式相结合

#运用策略模式与工厂模式相结合
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
    
测试：

    cashContext := new(CashContext)
    cashContext.CbxType = 赋值
    result = cashContext.GetResult(total)




