#简单工程模式
简单工厂模式是属于创建型模式，又叫做静态工厂方法（Static Factory Method）模式，但不属于23种设计模式之一。
简单工厂模式是由一个工厂对象决定创建出哪一种产品类的实例。
简单工厂模式是工厂模式家族中最简单实用的模式，可以理解为是不同工厂模式的一个特殊实现。

#需求
实现一个简单的计算器

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

#运用简单工程模式

    var a,b int
    
    type Calculate3 interface {
    GetResult() int
    }
    
    type Operation struct {
    a int
    b int
    }
    
    type OperationAdd struct {
    *Operation
    }
    
    func (a *OperationAdd)GetResult() int {
    return a.a + a.b
    }
    
    type OperationSub struct {
    *Operation
    }
    
    func (s *OperationSub)GetResult() int {
    return s.a - s.b
    }
    
    type OperationMul struct {
    *Operation
    }
    
    func (m *OperationMul)GetResult() int {
    return m.a * m.b
    }
    
    type OperationDiv struct {
    *Operation
    }
    
    func (a *OperationDiv)GetResult() int {
    if a.b == 0{
    return -1
    }
    return a.a / a.b
    }
    
    type OperationFactory struct {
    
    }
    
    //输入运算符，工厂就实例化出合适的对象
    //通过多态形式，返回父类的方式实现了计算器的结果
    func (f *OperationFactory)createOperate(operate string) Calculate3 {
    operation := &Operation{a:a,b:b}
    switch operate {
    case "+":
    return &OperationAdd{operation}
    case "-":
    return &OperationSub{operation}
    case "*":
    return &OperationMul{operation}
    case "/":
    return &OperationDiv{operation}
    default:
    return nil
    }
    }

相比与之前的方法，当需要修改已经写好的代码时，只需要修改其中一小部分而不会对其它运行良好的代码产生影响。
当需要增加新的运算方式时也是如此，只需要新增一个对象，实现GetResult（）方法即可，再维护一下createOperate（）。
体现松耦合的思想。

#执行测试
    var operate string
    fmt.Println("请输入数字A：")
    fmt.Scanf("%d",&a)
    fmt.Println("请选择运算符（+、-、*、/）：")
    fmt.Scanf("%s",&operate)
    fmt.Println("请输入数字B：")
    fmt.Scanf("%d",&b)
    factory := new(OperationFactory)
	calculate := factory.createOperate(operate)
	fmt.Println("计算结果为（-1表示有错）：",calculate.GetResult())

##执行结果
    请输入数字A：
    1
    请选择运算符（+、-、*、/）：
    +
    请输入数字B：
    2
    计算结果为（-1表示有错）： 3

#总结
1. 虽然go不是面向对象语言，但它可以通过组合方式实现面向对象的继承，简单来说就是新类型代理了其嵌入类型中的所有方法或者属性。
当外界调用新类型的方法或者属性时，Go的编译器首先会查看新类型有没有实现了这个方法或者属性，如果没有，就委托给新类型内部的
嵌入类型去实现这个方法或者调用属性去实现。这种组合方式主要是运用在interface和结构体中。
2. 从go运用简单工程模式这个小demo中，发现代码实现有点繁琐，不用这个简单工程模式而通过业务的封装方式写出来的代码反而没有那么繁琐，
这可能是go语言本身不太适合用这种方式来写吧。


