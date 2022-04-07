实体 | 说明
:---: | :---:
type LeiFeng interface | 雷锋（可理解基类）
type Person struct | 人，有三个学雷锋的动作Sweep、Wash、BuyRice
type Undergraduate struct | 毕业生，继承Person与其方法
type Volunteer struct | 志愿者，继承Person与其方法
type SimpleFactory struct | 简单工厂，用来来创建对象

###测试：
    func main()  {
    studentA := NewSimpleFactory("在校生")
    studentA.Wash()
    studentA.Sweep()
    studentA.BuyRice()
	volunteer := NewSimpleFactory("志愿者")
	volunteer.Wash()
	volunteer.Sweep()
	volunteer.BuyRice()
    }

###分析
1. 就是如果都是创建在校生的实例，NewSimpleFactory("在校生")这个方法就会重复很多遍，这就有些不够优雅
2. 如果要增加除在校生、志愿者等学雷锋的角色，就必须修改工厂方法的case分支条件，修改了原有的类，这不仅是对扩展开放了，也对修改开放了，违背了开放-封闭原则（软件实体可以扩展，但不可以修改）
