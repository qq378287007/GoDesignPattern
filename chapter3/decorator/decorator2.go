package main

import "fmt"

type Phone interface {
	GetPrice() float32
}

//装饰
type Decorator struct {
	Phone Phone
}

//装饰设置组件方法
func (d *Decorator) SetComponent(c Phone) {
	d.Phone = c
}

//装饰方法
func (d *Decorator) GetPrice() {
	if d.Phone != nil {
		d.Phone.GetPrice()
	}
}

//基础零件
type BaseParts struct {
}

//获取基础零件手机价格
func (p *BaseParts) GetPrice() float32 {
	return 2000
}

type Xiaomi struct {
	Decorator
}

//小米手机的价格
func (c *Xiaomi) GetPrice() float32 {
	phonePrice := c.Phone.GetPrice()
	return phonePrice + 1000
}

type IPhone struct {
	Decorator
}

//获取IPhone价格
func (c *IPhone) GetPrice() float32 {
	phonePrice := c.Phone.GetPrice()
	return phonePrice + 6000
}

func main() {
	//具体零件
	phone := &BaseParts{}
	fmt.Printf("基础零件的价格为：%f\n", phone.GetPrice())

	//定义添加IPhone手机
	iPhone := &IPhone{}
	iPhone.SetComponent(phone)
	fmt.Printf("苹果的价格为：%f\n", iPhone.GetPrice())

	//定义添加Xiaomi手机
	xiaomi := &Xiaomi{}
	xiaomi.SetComponent(phone)
	fmt.Printf("小米的价格为：%f\n", xiaomi.GetPrice())
}
