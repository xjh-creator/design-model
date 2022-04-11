###需求
通过建造小人的身体部位构造不同的小人

实体 | 说明
:---: | :---:
type ProductPerson struct | 产品
func (p ProductPerson)show() | 展示产品所有的部件
type Builder interface | 抽象建造者类
type ThinPersonBuilder struct | 具体建造者类，建造产品以及部件
func (t *ThinPersonBuilder)BuilderHead() | 建造部件头
func (t *ThinPersonBuilder)BuilderBody() | 建造部件身体
func (t *ThinPersonBuilder)BuilderLeftHand() | 建造部件左手
func (t *ThinPersonBuilder)BuilderRightHand() | 建造部件右手
func (t *ThinPersonBuilder)BuilderLeftLeg() | 建造部件左腿
func (t *ThinPersonBuilder)BuilderRightLeg() | 建造部件右腿
func (t ThinPersonBuilder)CreateProductPerson() ProductPerson | 创建产品
type Director struct | 指挥者类，用来指挥建造过程
func (d *Director)Construct() ProductPerson | 指挥者类用Construct来指挥建造过程
func NewDirector(builder Builder) *Director | 新建指挥者类


###测试：
    func main()  {
        thinPersonBuilder := ThinPersonBuilder{}
    
        director := NewDirector(&thinPersonBuilder)
        thinPerson := director.Construct()
        thinPerson.show()
    }

###分析：
1. 构建者模式将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示
2. 主要分为四个角色：（1）Product：具体产品 ,（2）ConcreteBuilder：具体建造者 ,（3）Builder：抽象建造者 ,（4）Director：指挥者
