实体 | 说明
:---: | :---:
type Component interface | 抽象构建，具体构建和抽象装饰类的基类，声明了在具体构建中实现的业务方法
type Person struct | 抽象构建的子类，用于定义具体的构建对象，装饰器可以给它增加额外的职责(方法)
func (p *Person)show() | 实现了在抽象构建中声明的方法
type Finery interface | Decorator，抽象装饰，也是抽象构建类的子类，用于给具体构建增加职责，但是具体职责在其子类中实现
type TShirts struct | ConcreteDecorator，具体装饰，抽象装饰的子类，负责向构建添加新的职责


