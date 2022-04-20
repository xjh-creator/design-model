###需求
展示不同手机品牌运行在它所在机子上的软件，此版本运用合成/聚合复用原则

实体 | 说明
:---: | :---:
type IHandsetBrand interface | 手机品牌
Run() | 方法，运行软件
SetHandsetSoft(soft IHandsetSoft) | 方法，设置手机的软件

实体 | 说明
:---: | :---:
type HandsetBrandM struct | 手机品牌M
NewHandsetBrandM() | 方法，初始化一个HandsetBrandM
SetHandsetSoft(soft IHandsetSoft) | 方法，设置手机的软件
Run() | 方法，运行手机的软件

实体 | 说明
:---: | :---:
type HandsetBrandN struct | 手机品牌M
NewHandsetBrandN() | 方法，初始化一个HandsetBrandN
SetHandsetSoft(soft IHandsetSoft) | 方法，设置手机的软件
Run() | 方法，运行手机的软件

实体 | 说明
:---: | :---:
type IHandsetSoft interface | 手机软件
Run() | 方法，运行软件
type HandsetGame struct | 游戏软件
NewHandsetGame() | 方法，初始化一个游戏软件
Run() | 方法，运行软件
type HandsetBrandAddressList struct | 通信软件
NewHandsetBrandAddressList() | 方法，初始化一个通信软件
Run() | 方法，运行软件



###测试：
    func main()  {
        var ab IHandsetBrand
        ab = NewHandsetBrandN()
    
        ab.SetHandsetSoft(NewHandsetGame())
        ab.Run()
    
        ab.SetHandsetSoft(NewHandsetBrandAddressList())
        ab.Run()
    
        ab = NewHandsetBrandM()
    
        ab.SetHandsetSoft(NewHandsetGame())
        ab.Run()
    
        ab.SetHandsetSoft(NewHandsetBrandAddressList())
        ab.Run()
    }

###分析：
1. 运用合成/聚合的方式，当需要增加手机品牌或者软件时，只需要增加一个类就行，不会对其它地方造成影响。
