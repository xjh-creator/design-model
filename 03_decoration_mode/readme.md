#装饰模式
装饰模式是为已有功能动态地添加更多功能的一种方式

##运用场景：
1. 系统需要新的功能的时候，向旧的类中添加新的代码，这些新的代码通常装饰了原有类的核心职责或主要行为。
2. 但新加入的代码只有某种特定情况下才会被用到，而旧的类代码是经常被用到的。
3. 装饰模式就是把要装饰的功能放在单独的类中，并让这个类来包装它所装饰的对象。需要用到时，就能有选择的把单独的类加入进来。

##代码结构：

实体 | 说明
:---: | :---:
type Component interface | 抽象构建，具体构建和抽象装饰类的基类，声明了在具体构建中实现的业务方法
type Person struct | 抽象构建的子类，用于定义具体的构建对象，装饰器可以给它增加额外的职责(方法)
func (p *Person)show() | 实现了在抽象构建中声明的方法
type Finery interface | Decorator，抽象装饰，也是抽象构建类的子类，用于给具体构建增加职责，但是具体职责在其子类中实现
type TShirts struct | ConcreteDecorator，具体装饰，抽象装饰的子类，负责向构建添加新的职责

##代码

    // 实现Component抽象类以及ConcreteComponent具体构建
    type Component interface {
        show()
    }
    
    //Person 是ConcreteComponent具体构建
    type Person struct {
        name string
    }
    
    func (p *Person)show()  {
        fmt.Println("装扮的",p.name)
    }
    
    //Finery 是 Decorator抽象装饰
    type Finery interface {
        Component
    }
    
    //TShirts 是ConcreteDecorator，具体装饰
    type TShirts struct {
        Finery
    }
    
    func (p *TShirts)show()  {
        fmt.Println("大T-shirt")
        p.Finery.show()
    }
    
    type BigTrouser struct {
        Finery
    }
    
    func (p *BigTrouser)show()  {
        fmt.Println("跨裤")
        p.Finery.show()
    }
    
    type BigSneakers struct {
        Finery
    }
    
    func (p *BigSneakers)show()  {
        fmt.Println("破球鞋")
        p.Finery.show()
    }
    
    type BigSuit struct {
        Finery
    }
    
    func (p *BigSuit)show()  {
        fmt.Println("西装")
        p.Finery.show()
    }
    
    type BigTie struct {
        Finery
    }
    
    func (p *BigTie)show()  {
        fmt.Println("领带")
        p.Finery.show()
    }
    
    type BigLeatherShoes struct {
        Finery
    }
    
    func (p *BigLeatherShoes)show()  {
        fmt.Println("皮鞋")
        p.Finery.show()
    }
    
    func  NewDecorator(t string,decorator Finery) Finery {
    switch t {
        case "TShirts":
        return &TShirts{decorator}
        case "BigTrouser":
        return &BigTrouser{decorator}
        case "BigSneakers":
        return &BigSneakers{decorator}
        case "BigSuit":
        return &BigSuit{decorator}
        case "BigTie":
        return &BigTie{decorator}
        case "BigLeatherShoes":
        return &BigLeatherShoes{decorator}
        default:
        return nil
        }
    }
    
    func main()  {
        component := &Person{"小明"}
        concreteDecorator1 := NewDecorator("TShirts",component)
        concreteDecorator1.show()
    
        concreteDecorator2 := NewDecorator("BigSuit",component)
        concreteDecorator2.show()
    }
