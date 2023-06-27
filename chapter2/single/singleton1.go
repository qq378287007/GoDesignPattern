package main

import (
	"fmt"
	"sync"
)

// 单例类
type singleton struct {
	value int
}

// 声明私有变量
var instance1 *singleton
var once sync.Once

func GetInstance1() *singleton {
	once.Do(func() {
		instance1 = new(singleton)
		fmt.Println("创建单个实例")
	})
	return instance1
}

var instance2 *singleton

// 获取单例对象
func GetInstance2() *singleton {
	if instance2 == nil {
		instance2 = new(singleton)
		fmt.Println("创建单个实例")
	}
	return instance2
}

var instance3 *singleton

// 声明锁对象
var mutex3 sync.Mutex

// 加锁保证协程安全
func GetInstance3() *singleton {
	mutex3.Lock()
	defer mutex3.Unlock()
	if instance3 == nil {
		instance3 = new(singleton)
		fmt.Println("创建单个实例")
	}
	return instance3
}

var instance3_2 *singleton
var mutex3_2 = &sync.Mutex{}

// 加锁保证协程安全
func GetInstance3_2() *singleton {
	if instance3_2 == nil {
		mutex3_2.Lock()
		defer mutex3_2.Unlock()
		if instance3_2 == nil {
			instance3_2 = new(singleton)
			fmt.Println("创建单个实例")
		} else {
			fmt.Println("已创建单个实例")
		}
	} else {
		fmt.Println("已创建单个实例")
	}
	return instance3_2
}

var instance4 *singleton

// 声明锁对象
var mutex4 sync.Mutex

// 当对象为空时，对对象加锁，当创建好对象后，获取对象时就不用加锁了
func GetInstance4() *singleton {
	if instance4 == nil {
		mutex4.Lock()
		if instance4 == nil {
			instance4 = new(singleton)
			fmt.Println("创建单个实例")
		}
		mutex4.Unlock()
	}
	return instance4
}

func main() {
	for i := 0; i < 3; i++ {
		go GetInstance1()
		go GetInstance2()
		go GetInstance3()
		go GetInstance3_2()
		go GetInstance4()
	}

	fmt.Scanln()
}
