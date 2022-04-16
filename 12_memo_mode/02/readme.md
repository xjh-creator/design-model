###需求
保存游戏角色的状态，在结束战斗后能够恢复状态

实体 | 说明
:---: | :---:
GameRole | 游戏角色
NewGameRole() | 方法，初始化一个游戏角色
StateDisplay() | 方法，状态展示
GetInitState() | 方法，获得初始状态
Fight() | 方法，与boss战斗
SaveState() | 方法，把状态保存到备忘录中
RecoveryState(memento *GameRoleMemento) | 方法，从备忘录中恢复之前保存的状态
StateDisplay() | 方法，展示
GameRoleMemento | 角色状态存储箱，保存游戏角色的状态
Caretaker | 角色状态管理者


###测试：
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

###备忘录角色：
1. Originator(发起人):负责创建一个备忘录Memento，用以记录当前时刻它的内部状态，并可以使用备忘录恢复内部状态。Originator可根据需要决定Memento存储Originator的那些内部状态
2. Memento(备忘录)：负责存储Originator对象的内部状态，并可防止Originator以外的其他对象访问备忘录Memento。
3. Caretaker(管理者)：负责保存好备忘录Memento，不能对备忘录的内容进行操作或检查。

###分析：
1. 保存的细节都封装到Memento中了，需要要改变细节也不用影响客户端了。
2. 备忘录模式适合功能复杂的，但需要维护或者纪录属性历史的类。
3. Originator可以根据保存的Memento信息还原到前一状态。
4. 有些对象需要把信息保存在对象外，但又必须对象自己来读取，这时可以用备忘录把复杂对象的内部信息对其他的对象屏蔽起来。
5. 当角色的状态改变的时候，有可能这个状态无效，这时候使用存储起来的备忘录将状态复原。

