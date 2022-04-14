###需求
表现员工工作的一天的不同时间段的不同状态

实体 | 说明
:---: | :---:
type Work struct | 工作状态，分类版
WriteProgram() | 

###测试：
    func main()  {
       emergencyProjects := &Work{}
        emergencyProjects.hour = 9
        emergencyProjects.WriteProgram()
        emergencyProjects.hour = 10
        emergencyProjects.WriteProgram()
        emergencyProjects.hour = 12
        emergencyProjects.WriteProgram()
        emergencyProjects.hour = 13
        emergencyProjects.WriteProgram()
        emergencyProjects.hour = 14
        emergencyProjects.WriteProgram()
        emergencyProjects.hour = 17
        emergencyProjects.WriteProgram()
    
        emergencyProjects.finish = false
    
        emergencyProjects.WriteProgram()
        emergencyProjects.hour = 19
        emergencyProjects.WriteProgram()
        emergencyProjects.hour = 22
        emergencyProjects.WriteProgram()
    }

###分析：
1. WriteProgram()方法过长了，有许多判断分支，意味着它责任过大，无论任何状态都要经过它来判断，违背了“单一职责原则”
2. 任何需要的改动都要去改动这个WriteProgram()方法，违背了“开放-封闭原则”
3. 解决方案：需要把分支想办法变成一个又一个的类，增加时不会影响其它类，状态变化又在各自的类中完成。


