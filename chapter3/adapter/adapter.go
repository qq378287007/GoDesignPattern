package main

import "fmt"

// 要适配的目标接口
type Target interface {
	Execute()
}

// 需要被适配的类
type Adaptee struct {
}

// 执行方法
func (a *Adaptee) SpecificExecute() {
	fmt.Println("最终执行的方法")
}

// 新接口Target的适配器，继承了Adaptee类
type Adapter struct {
	*Adaptee
}

// 实现Target 接口，同时继承了 Adaptee 类
func (a *Adapter) Execute() {
	a.SpecificExecute()
}

// 适配的目标接口
type ObjectTarget interface {
	Execute()
}

// 需要被适配的类
type ObjectAdaptee struct {
}

// 适配者类的方法
func (b *ObjectAdaptee) SpecificExecute() {
	fmt.Println("最终执行的方法")
}

//新接口Target的适配器类，通过关联Adaptee类来实现
type ObjectAdapter struct {
	Adaptee ObjectAdaptee
}

// 适配器类的方法
func (p *ObjectAdapter) Execute() {
	p.Adaptee.SpecificExecute()
}

func main() {
	//创建客户端
	var adapter Target = &Adapter{}
	adapter.Execute()

	var adapter2 ObjectTarget = &ObjectAdapter{}
	adapter2.Execute()
}
