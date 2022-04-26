###需求
为客户做网站，多个客户需求一致。简单版本：每个客户为它创建一个网站的实例

实体 | 说明
:---: | :---:
type WebSite struct | 网站类
NewWebSite(name string) | 方法，新建一个实例
func(u WebSite)Use() | 方法

###测试：
    func main()  {
        fx := NewWebSite("产品展示")
        fx.Use()
        fy := NewWebSite("产品展示")
        fy.Use()
        fz := NewWebSite("产品展示")
        fz.Use()
    
        fl := NewWebSite("博客")
        fl.Use()
        fm := NewWebSite("博客")
        fm.Use()
        fn := NewWebSite("博客")
        fn.Use()
    }

###分析：
1. 每个产品展示就需要一个实例，但它们本质是一样的代码。很多网站是共用一套代码，只是数据不同
2. 解决方案：采用实例共享

