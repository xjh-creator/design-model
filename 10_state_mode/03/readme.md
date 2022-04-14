###需求
表现员工工作的一天的不同时间段的不同状态

实体 | 说明
:---: | :---:
type Work struct | 环境类，维护一个ConcreteState子类的实例，这个实例定义当前的状态
SetHour() | 方法，设置当前时间
SetState() | 方法，设置当前状态
SetFinish() | 方法，设置任务是否完成
WriteProgram() | 方法，调用状态的方法，展示当前状态
type State interface | 状态基类
type ForenoonState struct | 状态子类，上午工作状态
WriteProgram(w Work) | 展示当前状态
type ForenoonState struct | 状态子类，上午工作状态
type NoonState struct | 状态子类，中午工作状态
type AfterNoonState struct | 状态子类，下午工作状态
type EveningState struct | 状态子类，晚上状态
type SleepingState struct | 状态子类，睡眠状态
type RestState struct | 状态子类，休息状态

###测试：
    func main()  {
        emergencyProjects := new(Work)
        tState := &NoonState{}
        emergencyProjects.SetState(tState)
        emergencyProjects.SetFinish(true)
        emergencyProjects.SetHour(15)
        emergencyProjects.WriteProgram()
    }

###分析：
1. 好处：将与特定状态相关的行为局部化，并且将不同状态的行为分割开来
2. 做法：将特定状态的相关行为都放入一个对象中，由于所有与状态相关的代码都存在于某个ConcreteState中，所以通过定义新的子类可以容易地增加心得状态和转换
3. 将各种状态转移逻辑分布到State的子类之间，来减少相互之间的依赖。
4. 目的：消除庞大的分支条件
5. 应用场景：当一个对象的行为取决于它的状态，并且它必须在运行时刻根据状态改变它的行为时，应用状态模式。


