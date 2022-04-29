###需求
男性与女性在面对同样发生的状态不同的表现。访问者模式代码

实体 | 说明
:---: | :---:
type Action interface | 抽象类，访问者vistor
GetConclusion(person IPerson) | 方法，元素Element
type Success struct | 具体类，具体访问者
type Failing struct| 具体类，具体访问者
type Amativeness struct | 具体类，具体访问者
type IPerson interface | 抽象类，Element
Accept(vistor Action) | 方法，Element通常要定义的方法，接收一个访问者作为参数
type Man struct | 具体类，具体Element
type Woman struct | 具体类，具体Element
type ObjectStructure struct | 对象结构，提供一个高层的接口允许访问者访问它的元素
Attach(element IPerson) | 方法，添加元素
Detach(element IPerson) | 方法，删除元素
Display(visitor Action) | 方法，展示访问者访问元素时展示的状态

###测试：
    func main()  {
        o := NewObjectStructure()
        o.Attach(NewMan())
        o.Attach(NewWonman())
    
        v1 := NewSuccess()
        o.Display(v1)
    
        v2 := NewFailing()
        o.Display(v2)
    
        v3 := NewAmativeness()
        o.Display(v3)
    }

###访问者模式定义：
表示一个作用于某对象结果中的各元素的操作。它使你可以在不改变各元素的类的前提上定义作用于这些元素的新操作。【DP】

###角色：
1. Visitor: 访问者，为对象结构（ObjectStructure）的ConcreteElement的每一个类声明一个Visit操作。
2. ConcreteVisitor:具体访问者，实现Visitor声明的操作，每个操作实现算法的一部分，而该算法片段乃是对应结构中对象的类。
3. Element: 元素，定义一个Accept操作，它以一个访问者为参数。
4. ConcreteElement: 具体元素，实现Accept操作。
5. ObjectStructure: 能够枚举它的元素，可以提供一个高层的接口以允许访问者访问它的元素。

###分析：
1. 目的：把处理从数据结构中分离出来。
2. 访问者模式适用于数据结构相对稳定的系统，如果不断有新的对象（Element）增加进来，就不适合适用。
3. 优点：增加新的操作容易，新的操作就意味着增加一个新的访问者（双分派的影响）。

