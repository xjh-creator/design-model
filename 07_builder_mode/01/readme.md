###需求
用画笔画不同小人的身体

实体 | 说明
:---: | :---:
type Graphics struct | 画板
DrawEllipse | 画头
DrawRectangle | 画身体
DrawLine | 画手足
type ThinPerson struct | 瘦子
Build | 画瘦子的身体
type FatPerson struct | 胖子
Build | 画胖子的身体

###测试：
    func main()  {
    thinPerson := &ThinPerson{Pen: "blue"}
    thinPerson.Build()
	fatPerson := &FatPerson{Pen: "red"}
	fatPerson.Build()
}

}

###分析：
1. 在新增小人的时候，因编程不注意会导致身体不健全的问题依旧存在
2. 最好的办法是规定要创建的小人必须有哪些部位
