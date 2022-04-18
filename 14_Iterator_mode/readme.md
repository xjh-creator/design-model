###需求
重现公交车售票员挨个向不同成乘客收取公交费

Iterator实体 | 说明
:---: | :---:
type Iterator interface | 迭代抽象类
First() string | 迭代方法，第一个
Next() string | 迭代方法，下一个
IsDone() bool | 迭代方法，判断是否迭代完
CurrentItem() string | 迭代方法，当前迭代对象

ConcreteIterator实体 | 说明
:---: | :---:
type ConcreteIterator struct | 具体迭代类
NewConcreteIterator(aggregate *ConcreteAggregate) | 方法，初始化一个ConcreteIterator
First() string | 迭代方法，第一个
Next() string | 迭代方法，下一个
IsDone() bool | 迭代方法，判断是否迭代完
CurrentItem() string | 迭代方法，当前迭代对象

Aggregate实体 | 说明
:---: | :---:
type Aggregate interface | 聚集抽象类
CreateIterator() Iterator | 方法，创建一个迭代器

ConcreteAggregate实体 | 说明
:---: | :---:
type ConcreteAggregate interface | 具体聚集类
CreateIterator() Iterator | 方法，创建一个迭代器



###测试：
    func main()  {
        bus := NewConcreteAggregate() //公交车上的乘客，聚集对象
	    bus.Items = append(bus.Items,"大鸟","小菜","行李","老外","公交内部员工","小偷") //添加对象

	    seller := NewConcreteIterator(bus) //售票员，迭代器对象
	    for !seller.IsDone(){
		    fmt.Println(seller.CurrentItem(),"请买票！")
		    seller.Next()
	    }
    }

###迭代器模式概念：
迭代器模式（Iterator）,提供一种方法顺序访问一个聚合对象中各个元素，而不暴露对象的内部表示。

###组合模式角色：
1. 抽象聚合类: 定义一个抽象的容器
2. 具体聚合类: 实现上面的抽象类，作为一个容器，用来存放元素，等待迭代
3. 抽象迭代器: 迭代器接口，每个容器下都有一个该迭代器接口的具体实现
4. 具体迭代器: 根据不同的容器，需要定义不同的具体迭代器，定义了游标移动的具体实现

###分析：
1. 迭代器模式很多语言就有封装，比如java中的foreach in
2. 在这里Iterator接口仿佛多余，直接使用用ConcreteIterator也能达到效果，其实这是因为遍历方式有多样，
   demo中的ConcreteIterator只是在顺序遍历，有时需要随机遍历、倒转遍历等。

