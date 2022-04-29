###需求
男性与女性在面对同样发生的状态不同的表现。面向对象版本。

实体 | 说明
:---: | :---:
type IPerson interface | 抽象类，人
GetConclusion() | 方法，面对不同状态时的表现
type Man struct | 具体类，男人
type Woman struct | 具体类，女人


###测试：
    func main()  {
        persons := make([]IPerson, 0)

        man1 := NewMan()
        man1.Action = "成功"
        persons = append(persons, man1)
        woman1 := NewWonman()
        woman1.Action = "成功"
        persons = append(persons, woman1)
    
        man2 := NewMan()
        man2.Action = "失败"
        persons = append(persons, man2)
        woman2 := NewWonman()
        woman2.Action = "失败"
        persons = append(persons, woman2)
    
        man3 := NewMan()
        man3.Action = "恋爱"
        persons = append(persons, man3)
        woman3 := NewWonman()
        woman3.Action = "恋爱"
        persons = append(persons, woman3)
    
        for _, v := range persons {
            v.GetConclusion()
        }
    }


