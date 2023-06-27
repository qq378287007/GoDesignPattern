package main

import (
	"fmt"
)

// 生成器接口
type Builder interface {
	Build()
}

// 产品
type Product struct {
	Built bool
}

// 具体生成器，用于构建产品的生成器
type ConcreteBuilder struct {
	result Product
}

// 初始化具体生成器对象
func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{result: Product{false}}
}

// 生成产品
func (b *ConcreteBuilder) Build() {
	b.result = Product{true}
}

// 返回在生成步骤中生成的产品
func (b *ConcreteBuilder) GetResult() Product {
	//return Product{true}
	return b.result
}

// 主管
type Director struct {
	builder Builder
}

// 初始化主管对象
func NewDirector(builder Builder) Director {
	return Director{builder}
}

// 通过一系列步骤生成产品
func (d *Director) Construct() {
	d.builder.Build()
}

func main() {
	concreteBuilder := NewConcreteBuilder()
	product := concreteBuilder.GetResult()
	fmt.Printf("product: %V\n", product)

	director := NewDirector(&concreteBuilder)
	director.Construct()
	product = concreteBuilder.GetResult()
	fmt.Printf("product: %V\n", product)

}
