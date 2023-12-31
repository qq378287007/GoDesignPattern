package main

import "fmt"

//享元接口
type Flyweight interface {
	Operation()
}

// 创建具体享元类，可以共享以支持大型有效的对象数量
type ConcreteFlyweight struct {
	intrinsicState string
}

//具体享元对象初始化
func (fw ConcreteFlyweight) Init(intrinsicState string) {
	fw.intrinsicState = intrinsicState
}

//具体享元对象的方法
func (fw ConcreteFlyweight) Operation(extrinsicState string) string {
	fmt.Println(fw.intrinsicState)
	if extrinsicState != "" {
		return extrinsicState
	}
	return "empty extrinsicState"
}

// 创建一个新的具体享元类
func NewConcreteFlyweight(state string) *ConcreteFlyweight {
	return &ConcreteFlyweight{state}
}

// 创建用于创建和存储享元的享元工厂类
type FlyweightFactory struct {
	pool map[string]*ConcreteFlyweight
}

// 创建一个新的享元工厂对象
func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{pool: make(map[string]*ConcreteFlyweight)}
}

// 获取或创建具体享元对象
func (f *FlyweightFactory) GetFlyweight(state string) *ConcreteFlyweight {
	flyweight := f.pool[state]
	if flyweight == nil {
		flyweight = NewConcreteFlyweight(state)
		f.pool[state] = flyweight
	}
	return flyweight
}

func main() {
	factory := NewFlyweightFactory()
	flyweight1 := factory.GetFlyweight("Barry")
	flyweight2 := factory.GetFlyweight("Shirdon")

	fmt.Println(flyweight1.Operation("ok"))
	fmt.Println(flyweight2.Operation("good"))
}
