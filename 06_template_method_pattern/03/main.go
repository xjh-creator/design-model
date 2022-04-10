package main

import "fmt"

//老师出的试卷
type TestPaper struct {
	answer1 string
	answer2 string
	answer3 string
}

func (t *TestPaper)TestQuestion1()  {
	fmt.Println("题目1，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	fmt.Println("答案：",t.answer1)
}

func (t *TestPaper)TestQuestion2()  {
	fmt.Println("题目2，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	fmt.Println("答案：",t.answer2)
}

func (t *TestPaper)TestQuestion3()  {
	fmt.Println("题目3，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	fmt.Println("答案：",t.answer3)
}

//学生甲抄的试卷
type TestPaperA struct {
	*TestPaper
}

//学生乙抄的试卷
type TestPaperB struct {
	*TestPaper
}


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

