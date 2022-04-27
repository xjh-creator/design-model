###需求
按照音谱内容弹奏音符音阶。

实体 | 说明
:---: | :---:
type Expression interface | 抽象表达式，
Interpret(context *PlayContext) | 方法，翻译音谱
Excute(key string,value string) | 方法，弹奏音符音阶
type Note struct | 音符，具体表达式
type Scale struct | 音阶，具体表达式
type PlayContext struct | 内容，要被表达式处理的内容

###测试：
    func main()  {
        context := NewPlayContext("")
        //音乐
        fmt.Println("上海滩")
        context.PlayText = "O2E1G2A3E1G1D2E2G1A3O3C1O2A1G1C2E1D3"
        var expression Expression
        for len(context.PlayText) > 0{
            str := string(context.PlayText[0])
            switch str {
                case "O":
                    expression = NewScale()
                default:
                    expression = NewNote()
            }
            expression.Interpret(context)
        }
    }

###定义：
解释器模式：给定一个语言，定义它的文法的一种表示，并定义一个解释器，这个解释器使用该表示来解释语言中的句子。

###角色：
Context:包含解释器之外的一些全局信息。
AbstractExpression:抽象表达式，也称抽象解释器，声明一个抽象的解释操作。
TerminalExpression:终结符表达式，实现与文法中的终结符相关联的解释操作。
NonTerminalExpression:非终结符表达式，为文法中的非终结符实现解释操作。

###分析：
1. 容易改变和扩展文法，因为该模式使用类来表示文法规则，可使用继承来改变或扩展文法
2. 容易实现文法，定义抽象语法树中各个节点的类的实现大体相同，这些类都易于直接编写。
3. 缺点：每个文法规则至少定义一个类，因此包含许多规则的文法可能难以管理和维护。不适用于文法非常复杂的情况。

