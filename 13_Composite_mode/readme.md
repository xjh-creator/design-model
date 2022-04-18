###需求
把总公司的各个部门的功能复用到子公司的部门

Component实体 | 说明
:---: | :---:
type ICompany interfac | 是组合中的对象声明接口
Add(company ICompany) | 方法，添加节点
Remove(company ICompany) | 方法，删除节点
Display(depth int) | 方法，展示分支
LineOfDuty() | 方法，展示部门职责

Composite实体 | 说明
:---: | :---:
type ConcreteCompany struct | Composite，定义有枝接节点行为，用来存储子部件
NewConcreteCompany(name string) | 方法，初始化一个ConcreteCompany
Add(company ICompany) | 方法，添加节点
Remove(company ICompany) | 方法，删除节点
Display(depth int) | 方法，展示分支
LineOfDuty() | 方法，展示部门职责

Leaf实体 | 说明
:---: | :---:
type HRDepartment struct | Leaf，表示叶子节点对象，叶子节点没有子节点
NewHRDepartment(name string) | 方法，初始化一个HRDepartment
Add(company ICompany) | 方法，叶子节点不实现此功能
Remove(company ICompany) | 方法，叶子节点不实现此功能
Display(depth int) | 方法，展示分支
LineOfDuty() | 方法，展示部门职责
type FinanceDepartment struct | Leaf，表示叶子节点对象，叶子节点没有子节点
NewFinanceDepartment(name string) | 方法，初始化一个FinanceDepartment
Add(company ICompany) | 方法，叶子节点不实现此功能
Remove(company ICompany) | 方法，叶子节点不实现此功能
Display(depth int) | 方法，展示分支
LineOfDuty() | 方法，展示部门职责


###测试：
    func main()  {
        root := NewConcreteCompany("北京总公司")
        root.Add(NewHRDepartment("总公司人力资源部"))
        root.Add(NewHRDepartment("总公司财务部"))
    
        comp := NewConcreteCompany("上海华东分公司")
        comp.Add(NewHRDepartment("上海华东分公司人力资源部"))
        comp.Add(NewHRDepartment("上海华东分公司财务部"))
        root.Add(comp)
    
        comp1 := NewConcreteCompany("南京办事处")
        comp1.Add(NewHRDepartment("南京办事处人力资源部"))
        comp1.Add(NewHRDepartment("南京办事处财务部"))
        root.Add(comp1)
    
        comp2 := NewConcreteCompany("杭州办事处")
        comp2.Add(NewHRDepartment("杭州办事处人力资源部"))
        comp2.Add(NewHRDepartment("杭州办事处财务部"))
        root.Add(comp2)
    
        fmt.Println("\n结构图：")
        root.Display(1)
    }

###组合模式概念：
组合模式(Composite):将对象组合成树形结构以以表示‘部分-整体’的层次结构。组合模式使得用户对单个对象和组合对象的使用具有一致性。

###组合模式角色：
1. Component:是组合中的对象声明接口，在适当的情况下，实现所有类共有接口的默认行为。声明一个接口用于访问和管理Component
2. Leaf:在组合中表示叶子节点对象，叶子节点没有子节点
3. Composite:定义有枝接节点行为，用来存储子部件，在Composite接口中实现与子部件有关操作
  

###分析：
1. 组合模式定义了具体部门的基本对象和分公司、办事处等组合对象的类层次结果，基本对象可以组合成更复杂的组合对象，而这个对象又可以继续被组合
，这样不断递归下去。
2. 组合模式可以让客户把组合对象当单个对象使用，不同关心是处理叶节点或者是一个组合组件。

