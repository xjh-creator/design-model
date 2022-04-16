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

//CreateMemento把状态保存到备忘录中
func (o *GameRole)SaveState() *GameRoleMemento {
	return &GameRoleMemento{
		vitality: o.vitality,
		attack: o.attack,
		defense: o.defense,
	}
}

//SetMemento恢复备忘录
func (o *GameRole)RecoveryState(memento *GameRoleMemento){
	o.vitality = memento.vitality
	o.attack = memento.attack
	o.defense = memento.defense
}

//StateDisplay状态显示
func (g GameRole)StateDisplay()  {
	fmt.Println("角色当前的状态：")
	fmt.Println("体力：",g.vitality)
	fmt.Println("攻击力：",g.attack)
	fmt.Println("防御力：",g.defense)
}

//GameRoleMemento角色状态存储箱
type GameRoleMemento struct {
	vitality int //生命力
	attack int //攻击力
	defense int //防御力
}
//Caretaker角色状态管理者
type Caretaker struct {
	memento *GameRoleMemento
}

func main()  {
	fmt.Println("大战boss前------------")
	xiaoming := NewGameRole()
	xiaoming.GetInitState()
	xiaoming.StateDisplay()

	//保存进度
	stateAdmin := &Caretaker{}
	stateAdmin.memento = xiaoming.SaveState()

	fmt.Println("大战Boss时，损耗严重------------")
	xiaoming.Fight()
	xiaoming.StateDisplay()

	fmt.Println("恢复之前的状态------------")
	xiaoming.RecoveryState(stateAdmin.memento)
    xiaoming.StateDisplay()
}
