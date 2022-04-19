###单例模式概念
保证一个类仅有一个实例，并提供一个访问它的全局访问点

###注意点
1. 某个类只能有一个实例
2. 这个类必须自行创建这个实例
3. 这个类必须自行向整个系统提供这个实例

###写法
1. 饿汉式：类在编译时创建自己的实例。由于是静态初始化，它是类一加载就实例化的对象，所以会占用资源
2. 懒汉式：类在运行时创建自己的实例。由于是在运行时才创建自己实例，会面临多线程访问的安全问题，需要处理才能保证安全。

###版本说明
####普通情况下单例模式
   
    var singleton *Singleton 

    type Singleton struct { 
    } 

    func GetInstance() *Singleton {
        if singleton == nil{
            singleton = &Singleton{}
        }
        return singleton
    }
####并发情况下的单例模式

    var singleton *Singleton
    var lock *sync.Mutex = &sync.Mutex{}
    
    type Singleton struct {
    
    }
    
    func GetInstance() *Singleton {
        lock.Lock()
        defer lock.Unlock()
        if singleton == nil{
            singleton = &Singleton{}
        }
        return singleton
    }
####上一种方式的改进，先判断实例是否存在再进行处理，不必每次都处理

    var singleton *Singleton
    var lock *sync.Mutex = &sync.Mutex{}
    
    type Singleton struct {
    
    }
    
    func GetInstance() *Singleton {
        if singleton == nil{
            lock.Lock()
            defer lock.Unlock()
            if singleton == nil{
                singleton = &Singleton{}
            }
        }
        return singleton
    }
####优雅的方式，使用Go独特的机制sync.Once,保证实例化代码只会执行一次

    var singleton *Singleton
    var once *sync.Once = &sync.Once{}
    
    type Singleton struct {
    
    }
    
    func GetInstance() *Singleton {
        once.Do(func() {
            singleton = &Singleton{}
        })
        return singleton
    }
