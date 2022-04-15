###需求
使用翻译替外籍球员翻译进攻与防守战术

实体 | 说明
:---: | :---:
type IPlayer interface | IPlayer（Target，目标角色接口），定义客户所需要的接口
Attack() | 方法，进攻战术
Defense() | 方法，防守战术
type Forwards struct | 前锋，正常角色，不需要通过适配器来调用方法
Attack() | 方法，进攻战术
Defense()t | 方法，防守战术
type Center struct | 中锋，正常角色，不需要通过适配器来调用方法
Attack() | 方法，进攻战术
Defense()t | 方法，防守战术
type Guards struct | 后卫，正常角色，不需要通过适配器来调用方法
Attack() | 方法，进攻战术
Defense()t | 方法，防守战术

实体 | 说明
:---: | :---:
type Translator struct | 适配器角色（Adapter，适配器）实现了客户调用的接口和另一接口
NewTranslator(iadaptee Iadaptee) | 方法，初始化一个适配器
Attack() | 方法，进攻战术
Defense()t | 方法，防守战术

实体 | 说明
:---: | :---:
type Iadaptee interface | 被转换角色（Adaptee）是需要被适配器转换成客户调用的接口
type ForeignCenter struct | 外籍中锋球员，需要翻译（适配器）
进攻() | 方法，进攻战术
防守() | 方法，防守战术
NewAdaptee(name string) | 方法，初始化一个被转换角色


###测试：
    func main()  {
        yaoming := NewAdaptee("yaoming")
        translator := NewTranslator(yaoming)
        translator.Attack()
        translator.Defense()
    }

###适配器模式
####简介
将一个接口转换成客户希望的另一个接口，使接口不兼容的那些类可以一起工作。
####组成
1. Target(目标接口角色): 定义了客户所需要的接口
2. Adapter(适配器角色): 实现了Target接口，并且还能调用另一个接口，起到适配的作用
3. Adaptee(被适配角色): 即被适配的角色，定义了一个已经存在的接口

###分析：
1. 要使用一个类时，如果它的方法与自己的要求不符合的时候，就需要考虑适配器模式。
2. 适配器模式是双方都不太容易修改的情况下才使用，接口能够重构就重构



