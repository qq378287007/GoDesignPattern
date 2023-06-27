package main

import "fmt"

type singleton struct {
}

var instance *singleton

func init() {
	//if instance == nil {
	instance = new(singleton)
	fmt.Println("创建单个实例")
	//}
}

// 提供获取实例的方法
func GetInstance() *singleton {
	return instance
}

func main() {
	for i := 0; i < 3; i++ {
		go GetInstance()
	}

	fmt.Scanln()
}
