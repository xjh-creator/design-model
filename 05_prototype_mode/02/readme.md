实体 | 说明
:---: | :---:
type PersonalInfo struct | 个人信息
type WorkExperience struct | 工作经历
type Resume struct | 个人简历，组合PersonalInfo，WorkExperience，也是原型类，有其实例来指定创建对象的种类，然后拷贝这个原型创建心得的对象
SetPersonalInfo | 个人简历方法
SetWorkExperience | 个人简历方法
Display | 个人简历方法
Clone | 克隆自身的方法

###注意：
1. go 语言中的传递都是值传递，传递一个对象，就会把对象拷贝一份传入函数中，传递一个指针，就会把指针拷贝一份传入进去。
2. 赋值的时候也是这样，res:=*e 就会把传递的 Example 对象拷贝一份，如果是 res:=e 的话，那么拷贝的就是对象的指针了.
3. 深拷贝和浅拷贝也可以这样理解，深拷贝就是拷贝整个对象，浅拷贝就是拷贝对象指针。