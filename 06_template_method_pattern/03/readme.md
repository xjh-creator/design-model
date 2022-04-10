###需求
两个同学抄写老师的题目并写出自己的答案

实体 | 说明
:---: | :---:
type TestPaper struct | 老师出的试卷
TestQuestion1 | 试题1
TestQuestion2 | 试题2
TestQuestion3 | 试题3
type TestPaperA struct | 学生甲抄的试卷
type TestPaperB struct | 学生乙抄的试卷

###测试：
    func main()  {
        fmt.Println("学生甲抄的试卷：")
        studentA := &TestPaperA{&TestPaper{answer1: "b",answer2: "c",answer3: "a"},}
        studentA.TestQuestion1()
        studentA.TestQuestion2()
        studentA.TestQuestion3()
    
        fmt.Println("学生乙抄的试卷：")
        studentB := &TestPaperB{&TestPaper{answer1: "c",answer2: "a",answer3: "a"},}
        studentB.TestQuestion1()
        studentB.TestQuestion2()
        studentB.TestQuestion3()

}
