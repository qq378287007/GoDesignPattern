package main

import "fmt"

//访问者接口
type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForRectangle(*Rectangle)
}

//形状接口
type Shape interface {
	GetType() string
	Accept(Visitor)
}

//圆形
type Circle struct {
	Radius float32
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

//矩形
type Rectangle struct {
	L float32
	B float32
}

func (t *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(t)
}

func (t *Rectangle) GetType() string {
	return "Rectangle"
}

//具体元素——正方形
type Square struct {
	Side float32
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}

//中点坐标计算器
type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) VisitForSquare(s *Square) {
	//...省略具体逻辑
	fmt.Println("计算正方形的中点坐标")
}

func (a *MiddleCoordinates) VisitForCircle(c *Circle) {
	//...省略具体逻辑
	fmt.Println("计算圆形的中点坐标")
}

func (a *MiddleCoordinates) VisitForRectangle(t *Rectangle) {
	//...省略具体逻辑
	fmt.Println("计算矩形的中点坐标")
}

//具体访问者——周长计算器
type PerimeterCalculator struct {
	perimeter float32
}

func (a *PerimeterCalculator) VisitForSquare(s *Square) {
	var perimeter float32
	perimeter = 4 * s.Side
	fmt.Printf("计算正方形的周长：%f \n", perimeter)
}

func (a *PerimeterCalculator) VisitForCircle(c *Circle) {
	var perimeter float32
	perimeter = 3.14 * 2 * c.Radius
	fmt.Printf("计算圆形的周长：%f \n", perimeter)
}

func (a *PerimeterCalculator) VisitForRectangle(r *Rectangle) {
	var perimeter float32
	perimeter = 2*r.B + 2*r.L
	fmt.Printf("计算矩形的周长：%f \n", perimeter)
}

func main() {
	//声明边长为1的正方形
	Square := &Square{Side: 1}
	//声明半径为6的圆形
	Circle := &Circle{Radius: 6}
	//声明长为8，宽为6的矩形
	Rectangle := &Rectangle{L: 8, B: 6}

	//声明周长计算器
	PerimeterCalculator := &PerimeterCalculator{}
	//正方形计算周长
	Square.Accept(PerimeterCalculator)
	//圆形计算周长
	Circle.Accept(PerimeterCalculator)
	//矩形计算周长
	Rectangle.Accept(PerimeterCalculator)
	//Square.GetType()

	fmt.Println()

	//声明中点坐标计算器
	MiddleCoordinates := &MiddleCoordinates{}
	//获取正方形中点坐标
	Square.Accept(MiddleCoordinates)
	//获取圆形中点坐标
	Circle.Accept(MiddleCoordinates)
	//获取矩形中点坐标
	Rectangle.Accept(MiddleCoordinates)
}
