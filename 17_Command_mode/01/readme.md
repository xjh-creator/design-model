###需求
演示路边烤羊肉串的实现。此版本使用紧耦合设计

实体 | 说明
:---: | :---:
type Barbecuer struct | 烤羊肉串者
NewBarbecuer() | 方法，初始化一个烤羊肉串者
BakeMutton() | 方法，烤羊肉串
BakeChickenWing() | 方法，烤鸡翅

###测试：
    func main()  {
        boy := NewBarbecuer()
        boy.BakeMutton()
        boy.BakeMutton()
        boy.BakeMutton()
        boy.BakeChickenWing()
        boy.BakeMutton()
        boy.BakeMutton()
        boy.BakeChickenWing()
    }


