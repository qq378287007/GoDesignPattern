package main

import "fmt"

//定义观察者接口
type Observer interface {
	Update(string)
	GetID() string
}

//定义用户模型
type User struct {
	Id string
}

//更新
func (c *User) Update(ProductName string) {
	fmt.Printf("发送邮件给用户： %s ，商品： %s 上架啦～ \n", c.Id, ProductName)
}

//获取编号
func (c *User) GetID() string {
	return c.Id
}

//主体类
type Subject interface {
	Register(Observer Observer)
	Deregister(Observer Observer)
	NotifyAll()
}

//商品类
type Product struct {
	ObserverList []Observer
	name         string
	inStock      bool
}

//新增商品
func NewProduct(name string) *Product {
	return &Product{name: name}
}

//更新库存状态
func (i *Product) UpdateAvailability() {
	fmt.Printf("商品 %s 现在上架了～\n", i.name)
	i.inStock = true
	i.NotifyAll()
}

//注册到观察者列表
func (i *Product) Register(o Observer) {
	i.ObserverList = append(i.ObserverList, o)
}

//从观察者列表删除
func (i *Product) Deregister(o Observer) {
	i.ObserverList = RemoveFromslice(i.ObserverList, o)
}

//通知所有用户
func (i *Product) NotifyAll() {
	for _, Observer := range i.ObserverList {
		Observer.Update(i.name)
	}
}

func RemoveFromslice(ObserverList []Observer, ObserverToRemove Observer) []Observer {
	ObserverListLength := len(ObserverList)
	for i, Observer := range ObserverList {
		if ObserverToRemove.GetID() == Observer.GetID() {
			ObserverList[ObserverListLength-1], ObserverList[i] = ObserverList[i], ObserverList[ObserverListLength-1]
			return ObserverList[:ObserverListLength-1]
		}
	}
	return ObserverList
}

func main() {
	//声明商品对象
	bookProduct := NewProduct("《Go语言高级开发与实战》")

	ObserverFirst := &User{Id: "shirdonliao@gmail.com"}
	ObserverSecond := &User{Id: "barry@gmail.com"}

	//通知第一个用户
	bookProduct.Register(ObserverFirst)
	//通知第二个用户
	bookProduct.Register(ObserverSecond)

	//更新库存
	bookProduct.UpdateAvailability()
}
