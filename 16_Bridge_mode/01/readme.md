###需求
重现公交车售票员挨个向不同成乘客收取公交费

实体 | 说明
:---: | :---:
type IHandsetBrand interface | 手机品牌
Run() | 方法，运行软件


实体 | 说明
:---: | :---:
type HandsetBrandM struct | 手机品牌M
NewHandsetBrandM() | 方法，初始化一个HandsetBrandM
type HandsetBrandMGame struct | 手机品牌M的游戏软件
Run() | 方法，运行软件
type HandsetBrandMAddressList struct | 手机品牌M的通讯软件
Run() | 方法，运行软件

实体 | 说明
:---: | :---:
type HandsetBrandN struct | 手机品牌N
NewHandsetBrandN() | 方法，初始化一个HandsetBrandN
type HandsetBrandNGame struct | 手机品牌N的游戏软件
Run() | 方法，运行软件
type HandsetBrandNAddressList struct | 手机品牌N的通讯软件
Run() | 方法，运行软件



###测试：
    func main()  {
       var ab IHandsetBrand
       ab = HandsetBrandMAddressList{NewHandsetBrandM()}
       ab.Run()
   
       ab = HandsetBrandMGame{NewHandsetBrandM(),}
       ab.Run()
   
       ab = HandsetBrandNAddressList{NewHandsetBrandN()}
       ab.Run()
   
       ab = HandsetBrandNGame{NewHandsetBrandN()}
       ab.Run()
    }

###分析：
1. 在增加功能时，一般思路是在每个手机具体品牌下面新增一个软件功能，而增加一个手机品牌就是在手机品牌下面新增一个具体品牌类，
总结就是不断使用继承增加新类。
2. 对象的继承关系是在编译是就定义好，无法在运行时改变从父类继承的实现。子类的实现与父类有非常紧密的依赖关系，父类实现中任何变化都会导致子类发生变化，
如果继承下来的实现不适合解决新问题，则父类必须重写或者被更适合的类替代。这种依赖关系限制了灵活性并且最终限制了复用性。
3. 改进思路：使用合成/聚合复用原则（尽量使用合成/聚合，尽量不要使用类继承）
4. 合成：强拥有关系（部分与整体的关系，生命周期一样）。如两个翅膀合成大雁。
5. 聚合：弱拥有关系（A包含B,但B可以不是A的一部分）。如多只大雁聚合成大雁群。
6. 作用：帮助每个类被封装，并被集中到某个任务上，这样类和继承会保持较小规模，并且不会成为不可控的庞然大物（如手机品牌、具体手机品牌、具体手机品牌的软件的例子）

