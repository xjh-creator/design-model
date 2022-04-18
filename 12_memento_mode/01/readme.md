###需求
保存游戏角色的状态，在结束战斗后能够恢复状态

实体 | 说明
:---: | :---:
GameRole | 游戏角色
NewGameRole() | 方法，初始化一个游戏角色
StateDisplay() | 方法，状态展示
GetInitState() | 方法，获得初始状态
Fight() | 方法，与boss战斗

###测试：
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

###分析：
1. 问题：这种恢复体力的写法把游戏角色的细节暴露给了客户端，客户端的职责太大，需要知道角色属性并对它进行备份。
###暴露的细节部分
      backup.vitality = xiaoming.vitality
      backup.attack = xiaoming.attack
      backup.defense = xiaoming.defense


