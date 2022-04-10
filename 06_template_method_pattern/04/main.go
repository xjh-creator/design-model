package main

import "fmt"

type ITestPaper interface {
	TestQuestion1()
	TestQuestion2()
	TestQuestion3()
	Answer1() string
	Answer2() string
	Answer3() string
}

//老师出的试卷
type TestPaper struct {
	Specific ITestPaper
}

func (t *TestPaper)TestQuestion1()  {
	fmt.Println("题目1，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	if t.Specific != nil{
		fmt.Println("答案：",t.Specific.Answer1())
		return
	}
	fmt.Println("答案：",t.Answer1())
}

func (t *TestPaper)TestQuestion2()  {
	fmt.Println("题目2，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	if t.Specific != nil{
		fmt.Println("答案：",t.Specific.Answer2())
		return
	}
	fmt.Println("答案：",t.Answer2())
}

func (t *TestPaper)TestQuestion3()  {
	fmt.Println("题目3，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	if t.Specific != nil{
		fmt.Println("答案：",t.Specific.Answer3())
		return
	}
	fmt.Println("答案：",t.Answer3())
}

func (t *TestPaper)Answer1() string {
	return ""
}

func (t *TestPaper)Answer2() string {
	return ""
}

func (t *TestPaper)Answer3() string {
	return ""
}

//学生甲抄的试卷
type TestPaperA struct {
	*TestPaper
}

func (t *TestPaperA)Answer1() string {
	return "b"
}

func (t *TestPaperA)Answer2() string {
	return "c"
}

func (t *TestPaperA)Answer3() string {
	return "a"
}

//学生乙抄的试卷
type TestPaperB struct {
	*TestPaper
}

func (t *TestPaperB)Answer1() string {
	return "c"
}

func (t *TestPaperB)Answer2() string {
	return "a"
}

func (t *TestPaperB)Answer3() string {
	return "a"
}


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


