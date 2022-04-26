###需求
为客户做网站，多个客户需求一致。享元模式：每个客户的需求如果一致，即他们共享一个网站的实例

实体 | 说明
:---: | :---:
type WebSite interface| 网站抽象类，基类，也是享员类的基类
Use() | 方法，使用网站
type ConcreteWebSite struct | 具体网站，有一个字段Name，享元类
Use() | 方法，使用网站
type WebSiteFactory struct | 网站工厂，用来创建并管理享元对象
GetWebSiteCategory(key string) | 方法，得到网站，如果网站实例不存在即新建
GetWebSiteCount() | 方法，得到实例的个数

###测试：
    func main()  {
        f := NewWebSiteFactory()

        fx := f.GetWebSiteCategory("产品展示")
        fx.Use()
    
        fy := f.GetWebSiteCategory("产品展示")
        fy.Use()
    
        fz := f.GetWebSiteCategory("产品展示")
        fz.Use()
    
        f1 := f.GetWebSiteCategory("博客")
        f1.Use()
    
        fm := f.GetWebSiteCategory("博客")
        fm.Use()
    
        fn := f.GetWebSiteCategory("博客")
        fn.Use()
    
        fmt.Println("网站分类总数为",f.GetWebSiteCount())
    }

###分析：
1. 实现了享元模式共享对象的目的，不管新建多少个网站，如果名字相同，即共享一个实例

