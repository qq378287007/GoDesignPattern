package main

import "fmt"

// 原型接口
type Prototype interface {
	GetName() string
	Clone() Prototype
}

// 具体原型类
type ConcretePrototype struct {
	Name string
}

// 返回具体原型的名称
func (p *ConcretePrototype) GetName() string {
	return p.Name
}

// Clone 创建一个ConcretePrototype类的克隆新实例
func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{p.Name}
}

func main() {
	cp := &ConcretePrototype{Name: "Shirdon"}
	res := cp.GetName()
	fmt.Println(res)

	cp2 := cp.Clone()
	res2 := cp2.GetName()
	fmt.Println(res2)
}
