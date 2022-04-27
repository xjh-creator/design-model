package main

import "fmt"

//Expression表达式
type Expression interface {
	Interpret(context *PlayContext)
	Excute(key string,value string)
}

//Note音符类TerminalExpression
type Note struct {

}

func NewNote() *Note {
	return &Note{}
}

func (n *Note)Excute(key string,value string)  {
	note := ""
	switch key {
		case "C":note= "1"
		case "D":note= "2"
		case "E":note= "3"
		case "F":note= "4"
		case "G":note= "5"
		case "A":note= "6"
		case "B":note= "7"
	}
	fmt.Println("音符",note)
}

func (n *Note)Interpret(context *PlayContext)  {
	if len(context.PlayText) == 0{
		return
	}else{
		playText := []byte(context.PlayText)
		playKey := string(playText[0])
		playValue := string(playText[1])
		context.PlayText = string(playText[2:])
		n.Excute(playKey,playValue)
	}
}

//Scale音阶类TerminalExpression
type Scale struct {

}

func NewScale() *Scale {
	return &Scale{}
}

func (s *Scale)Excute(key string,value string)  {
	scale := ""
	switch value {
	case "1":scale = "低音"
	case "2":scale = "中音"
	case "3":scale = "高音"
	}
	fmt.Println("音阶",scale)
}

func (s *Scale)Interpret(context *PlayContext)  {
	if len(context.PlayText) == 0{
		return
	}else{
		playText := []byte(context.PlayText)
		playKey := string(playText[0])
		playValue := string(playText[1])
		context.PlayText = string(playText[2:])
		s.Excute(playKey,playValue)
	}
}

type PlayContext struct {
	PlayText string
}

func NewPlayContext(text string) *PlayContext {
	return &PlayContext{PlayText: text}
}

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


