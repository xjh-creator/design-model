###单例模式概念
保证一个类仅有一个实例，并提供一个访问它的全局访问点

###注意点
1. 某个类只能有一个实例
2. 这个类必须自行创建这个实例
3. 这个类必须自行向整个系统提供这个实例

###写法
1. 饿汉式：类在编译时创建自己的实例。由于是静态初始化，它是类一加载就实例化的对象，所以会占用资源。简单理解就是实例已经提前创建好了，
GetInstance()只是获取实例返回。
2. 懒汉式：类在运行时创建自己的实例。由于是在运行时才创建自己实例，会面临多线程访问的安全问题，需要处理才能保证安全。

###版本说明
####饿汉式单例模式

    func init() {
        s = &singleton{}
    }
    
    var s *singleton
    
    type singleton struct {
    }
    
    func GetInstance() *singleton {
        return s
    }
    
    func main() {
        a1 := GetInstance()
        a2 := GetInstance()
        fmt.Println(a1 == a2) //结果为true
    }
####普通情况下单例模式

    var s *singleton
    
    type singleton struct {
    }
    
    func GetInstance() *singleton {
    if s == nil {
        s = &singleton{}
    }
        return s
    }
####并发情况下的单例模式

    var (
        s    *singleton
        lock = &sync.Mutex{}
    )
    
    type singleton struct {
    }
    
    func GetInstance() *singleton {
        lock.Lock()
        defer lock.Unlock()
        if s == nil {
            s = &singleton{}
        }
        return s
    }
####上一种方式的改进，先判断实例是否存在再进行处理，不必每次都处理

    var (
        s    *singleton
        lock = &sync.Mutex{}
    )
    
    type singleton struct {
    }
    
    func GetInstance() *singleton {
        if s == nil {
            lock.Lock()
            defer lock.Unlock()
            if s == nil {
                s = &singleton{}
            }
        }
        return s
    }

####优雅的方式，使用Go独特的机制sync.Once,保证实例化代码只会执行一次

    var (
	s    *singleton
	once = &sync.Once{}
    )

    type singleton struct {
    }
    
    func GetInstance() *singleton {
        once.Do(func() {
            s = &singleton{}
        })
        return s
    }
