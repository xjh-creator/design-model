###需求
展示不同手机品牌运行在它所在机子上的软件，此版本运用桥接模式，其实桥接模式就是运用了合成/聚合原则

实体 | 说明
:---: | :---:
type IHandsetBrand interface | 手机品牌
Run() | 方法，运行软件
SetHandsetSoft(soft IHandsetSoft) | 方法，设置手机的软件
NewIHandsetBrand(handsetBrand string,handsetSoft IHandsetSoft) | 简单工厂方法，创建一个手机品牌

实体 | 说明
:---: | :---:
type HandsetBrandM struct | 手机品牌M
NewHandsetBrandM() | 方法，初始化一个HandsetBrandM
SetHandsetSoft(soft IHandsetSoft) | 方法，设置手机的软件
Run() | 方法，运行手机的软件

实体 | 说明
:---: | :---:
type HandsetBrandN struct | 手机品牌M
NewHandsetBrandN() | 方法，初始化一个HandsetBrandN
SetHandsetSoft(soft IHandsetSoft) | 方法，设置手机的软件
Run() | 方法，运行手机的软件

实体 | 说明
:---: | :---:
type IHandsetSoft interface | 手机软件
Run() | 方法，运行软件
type HandsetGame struct | 游戏软件
NewHandsetGame() | 方法，初始化一个游戏软件
Run() | 方法，运行软件
type HandsetBrandAddressList struct | 通信软件
NewHandsetBrandAddressList() | 方法，初始化一个通信软件
Run() | 方法，运行软件



###测试：
    func main()  {
        handsetBrandM := NewIHandsetBrand("M",NewHandsetGame())
        handsetBrandM.Run()
        handsetBrandM.SetHandsetSoft(NewHandsetBrandAddressList())
        handsetBrandM.Run()
    
        handsetBrandN := NewIHandsetBrand("N",NewHandsetGame())
        handsetBrandN.Run()
        handsetBrandN.SetHandsetSoft(NewHandsetBrandAddressList())
        handsetBrandN.Run()
    }

###桥接模式：
1. 桥接模式(Bridge)，将抽象部分与它的实现部分分离，使它们都可以独立的变化。
2. 桥接模式主要使用抽象关联取代传统的多重继承，将类之间的静态继承关系转换为动态的对象组合关系，使得系统更加灵活，并易于扩展，同时有效控制了系统中类的个数。

###主要概念：
1. Abstraction（抽象类，图中的BrushPen接口）：用于定义抽象类的接口，它与Implementor之间具有关联关系，它既可以包含抽象业务方法， 也可以包含具体业务方法。
2. RefinedAbstratction（扩充抽象类，图中的BigBrushPen和SmallBrushPen）：扩充由Abstraction定义的接口，通常情况下他不再是抽象类而是具体类，实现了在Abstraction中声明的抽象业务方法，在RefinedAbstraction中可以调用在Implementor中定义的业务方法。
3. Implementor（实现类接口，图中的Color接口）：定义实现类的接口，一般而言，它不与Abstraction的接口一致。它只提供基本操作，而Abstraction定义的接口可能会做更多更复杂的操作。
4. ConcreteImplementor（具体实现类，图中的Red、Green、Yellow）：具体实现Implementor接口，在不同的ConcreteImplementor中提供基本操作的不同实现，在程序运行时，ConcreteImplentor将替换其父类对象，提供给抽象类具体的业务操作方法。

###总结
在面向对象中，能用合成/聚合解决的就不要用继承。