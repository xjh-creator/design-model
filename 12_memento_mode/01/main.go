package main

import "fmt"

//GameRole游戏角色
type GameRole struct {
	vitality int //生命力
	attack int //攻击力
	defense int //防御力
}

func NewGameRole() *GameRole {
	return &GameRole{}
}

//StateDisplay状态显示
func (g GameRole)StateDisplay()  {
	fmt.Println("角色当前的状态：")
	fmt.Println("体力：",g.vitality)
	fmt.Println("攻击力：",g.attack)
	fmt.Println("防御力：",g.defense)
}
//GetInitState获得初始状态
func (g *GameRole)GetInitState()  {
	g.vitality = 100
	g.attack = 100
	g.defense = 100
}

//Fight与boss战斗
func (g *GameRole)Fight()  {
	g.vitality = 0
	g.attack = 0
	g.defense = 0
}

func main()  {
	fmt.Println("大战boss前------------")
	xiaoming := NewGameRole()
	xiaoming.GetInitState()
	xiaoming.StateDisplay()

	//保存进度
	backup := NewGameRole()
	backup.vitality = xiaoming.vitality
	backup.attack = xiaoming.attack
	backup.defense = xiaoming.defense

	fmt.Println("大战Boss时，损耗严重------------")
	xiaoming.Fight()
	xiaoming.StateDisplay()

	fmt.Println("恢复之前的状态------------")
	xiaoming.vitality = backup.vitality
	xiaoming.attack = backup.attack
	xiaoming.defense = backup.defense

	xiaoming.StateDisplay()
}
