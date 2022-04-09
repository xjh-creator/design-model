###需求
两个同学抄写老师的题目并写出自己的答案

实体 | 说明
:---: | :---:
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
1. 每个实体非常相似，除了甲乙同学自己写的答案不同，实体，试题的题目都一样
2. 如果要改需求，如试题的题目，两个试题都要改
3. 这种最基本的写法容易写错，又难以维护。