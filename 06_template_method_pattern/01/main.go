package main

import "fmt"

//学生甲抄的试卷
type TestPaperA struct {

}

func (t *TestPaperA)TestQuestion1()  {
	fmt.Println("题目1，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	fmt.Println("答案：b")
}

func (t *TestPaperA)TestQuestion2()  {
	fmt.Println("题目2，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	fmt.Println("答案：a")
}

func (t *TestPaperA)TestQuestion3()  {
	fmt.Println("题目3，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	fmt.Println("答案：c")
}

//学生乙抄的试卷
type TestPaperB struct {

}

func (t *TestPaperB)TestQuestion1()  {
	fmt.Println("题目1，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	fmt.Println("答案：d")
}

func (t *TestPaperB)TestQuestion2()  {
	fmt.Println("题目2，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	fmt.Println("答案：b")
}

func (t *TestPaperB)TestQuestion3()  {
	fmt.Println("题目3，答案是【】。a.答案1 b.答案2 c.答案3 d.答案4")
	fmt.Println("答案：a")
}

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

