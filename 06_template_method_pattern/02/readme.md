###需求
两个同学抄写老师的题目并写出自己的答案

实体 | 说明
:---: | :---:
type TestPaper struct | 老师出的试卷
TestQuestion1 | 试题1
TestQuestion2 | 试题2
TestQuestion3 | 试题3
type TestPaperA struct | 学生甲抄的试卷
TestQuestion1 | 试题1
TestQuestion2 | 试题2
TestQuestion3 | 试题3
type TestPaperB struct | 学生乙抄的试卷
TestQuestion1 | 试题1
TestQuestion2 | 试题2
TestQuestion3 | 试题3

###测试：
    func main()  {
	fmt.Println("学生甲抄的试卷：")
	studentA := &TestPaperA{}
	studentA.TestQuestion1()
	studentA.TestQuestion2()
	studentA.TestQuestion3()

	fmt.Println("学生乙抄的试卷：")
	studentB := &TestPaperB{}
	studentB.TestQuestion1()
	studentB.TestQuestion2()
	studentB.TestQuestion3()

}

###分析：
1. 这是初步的泛化，但还有有很多重复的代码
2. 甲乙两个同学的代码除了答案不同之外，其他都是重复的
3. 用到继承的机制的话，重复的代码都应该要上升到父类去，子类中就不需要重复的代码