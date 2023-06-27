package main

import "fmt"

type Car struct {
	SeatsType  string
	EngineType string
	Number     int
}

//生成器接口
type InterfaceBuilder interface {
	SetSeatsType()
	SetEngineType()
	SetNumber()
	GetCar() Car
}

//MPV生成器
type MpvBuilder struct {
	SeatsType  string
	EngineType string
	Number     int
}

func NewMpvBuilder() *MpvBuilder {
	return &MpvBuilder{}
}

func (b *MpvBuilder) SetSeatsType() {
	b.SeatsType = "MPV型座椅"
}

func (b *MpvBuilder) SetEngineType() {
	b.EngineType = "MPV型引擎"
}

func (b *MpvBuilder) SetNumber() {
	b.Number = 8
}

func (b *MpvBuilder) GetCar() Car {
	return Car{
		EngineType: b.EngineType,
		SeatsType:  b.SeatsType,
		Number:     b.Number,
	}
}

//SUV生成器
type SuvBuilder struct {
	SeatsType  string
	EngineType string
	Number     int
}

func newSuvBuilder() *SuvBuilder {
	return &SuvBuilder{}
}

func (b *SuvBuilder) SetSeatsType() {
	b.SeatsType = "SUV型座椅"
}

func (b *SuvBuilder) SetEngineType() {
	b.EngineType = "SUV型引擎"
}

func (b *SuvBuilder) SetNumber() {
	b.Number = 6
}

func (b *SuvBuilder) GetCar() Car {
	return Car{
		EngineType: b.EngineType,
		SeatsType:  b.SeatsType,
		Number:     b.Number,
	}
}

//获取生成器
func GetBuilder(BuilderType string) InterfaceBuilder {
	if BuilderType == "mpv" {
		return &MpvBuilder{}
	}

	if BuilderType == "suv" {
		return &SuvBuilder{}
	}
	return nil
}

//主管类型
type Director struct {
	Builder InterfaceBuilder
}

func NewDirector(b InterfaceBuilder) *Director {
	return &Director{
		Builder: b,
	}
}

func (d *Director) SetBuilder(b InterfaceBuilder) {
	d.Builder = b
}

func (d *Director) BuildCar() Car {
	d.Builder.SetEngineType()
	d.Builder.SetSeatsType()
	d.Builder.SetNumber()
	return d.Builder.GetCar()
}

func main() {
	//声明MPV生成器对象
	MpvBuilder := GetBuilder("mpv")
	//声明主管对象
	Director := NewDirector(MpvBuilder)
	//生产MPV类型汽车
	mpvCar := Director.BuildCar()

	fmt.Printf("MPV类型引擎: %s\n", mpvCar.EngineType)
	fmt.Printf("MPV类型座椅: %s\n", mpvCar.SeatsType)
	fmt.Printf("MPV类型数量: %d\n", mpvCar.Number)

	//声明SUV生成器对象
	SuvBuilder := GetBuilder("suv")
	//设置生成器对象
	Director.SetBuilder(SuvBuilder)
	//生产SUV类型汽车
	suvCar := Director.BuildCar()

	fmt.Printf("\nSUV类型引擎: %s\n", suvCar.EngineType)
	fmt.Printf("SUV类型座椅: %s\n", suvCar.SeatsType)
	fmt.Printf("SUV类型数量: %d\n", suvCar.Number)
}
