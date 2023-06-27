package main

import "fmt"

//备忘录类状态
type Memento struct {
	State string
}

//获取设置好的状态
func (m *Memento) GetSavedState() string {
	return m.State
}

//负责人类
type Caretaker struct {
	MementoArray []*Memento
}

//添加备忘录
func (c *Caretaker) AddMemento(m *Memento) {
	c.MementoArray = append(c.MementoArray, m)
}

//获取备忘录
func (c *Caretaker) GetMemento(index int) *Memento {
	return c.MementoArray[index]
}

//原发器类
type Originator struct {
	State string
}

//创建备忘录
func (e *Originator) CreateMemento() *Memento {
	return &Memento{State: e.State}
}

//恢复原发器对象的状态
func (e *Originator) RestoreMemento(m *Memento) {
	e.State = m.GetSavedState()
}

//设置状态
func (e *Originator) SetState(State string) {
	e.State = State
}

//获取状态
func (e *Originator) GetState() string {
	return e.State
}

func main() {
	//声明负责人对象
	Caretaker := &Caretaker{MementoArray: make([]*Memento, 0)}

	//声明原发器对象
	Originator := new(Originator)
	Originator.State = "One"
	fmt.Printf("Originator 当前状态: %s\n", Originator.GetState())

	//添加备忘录
	Caretaker.AddMemento(Originator.CreateMemento())

	Originator.SetState("Two")
	fmt.Printf("Originator 当前状态: %s\n", Originator.GetState())

	//添加备忘录
	Caretaker.AddMemento(Originator.CreateMemento())

	Originator.SetState("Three")
	fmt.Printf("Originator 当前状态: %s\n", Originator.GetState())

	//添加备忘录
	Caretaker.AddMemento(Originator.CreateMemento())

	//恢复原发器对象的状态
	Originator.RestoreMemento(Caretaker.GetMemento(1))
	fmt.Printf("恢复到状态: %s\n", Originator.GetState())

	//恢复原发器对象的状态
	Originator.RestoreMemento(Caretaker.GetMemento(0))
	fmt.Printf("恢复到状态: %s\n", Originator.GetState())
}
