###需求
两个同学抄写老师的题目并写出自己的答案

###模板方法定义
模板方法模式，定义一个操作中的算法的骨架，而将一些步骤延迟到子类中，模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤

实体 | 说明
:---: | :---:
type ITestPaper interface | 接口
type TestPaper struct | 超类，老师出的试卷
TestQuestion1 | 试题1
TestQuestion2 | 试题2
TestQuestion3 | 试题3
Answer1 | 试题1的答案
Answer2 | 试题2的答案
Answer3 | 试题3的答案
type TestPaperA struct | 学生甲抄的试卷
Answer1 | 甲同学试题1的答案
Answer2 | 甲同学试题2的答案
Answer3 | 甲同学试题3的答案
type TestPaperB struct | 学生乙抄的试卷
Answer1 | 乙同学试题1的答案
Answer2 | 乙同学试题2的答案
Answer3 | 乙同学试题3的答案

###测试：
    func main()  {
        fmt.Println("学生甲抄的试卷：")
        studentA := &TestPaper{&TestPaperA{}}
        studentA.TestQuestion1()
        studentA.TestQuestion2()
        studentA.TestQuestion3()
    
        fmt.Println("学生乙抄的试卷：")
        studentB := &TestPaper{&TestPaperB{}}
        studentB.TestQuestion1()
        studentB.TestQuestion2()
        studentB.TestQuestion3()

}

###分析：
1. 模板方法模式就是把不变行为放到超类，去除子类中的重复代码来体现优势
2. 模板方法模式就是提供了一个很好的代码复用平台，适合代码运行过程在高层次运行相同，但有些步骤可能不同

###与Java的比较
1. golang:通过组合的方式 来实现java中类似的效果。
2. 一个接口，包含了java抽象类中所有的方法定义
3. 一个结构体，封装java抽象类中的非抽象方法，并持有一个接口引用对象
4. 具体的实现，匿名组合上述结构体，实现接口中未实现的方法（即java抽象类中的抽象方法）
5. 在调用结构体对象中的非抽象方法之前，需要将具体子类对象赋给它的接口引用。

