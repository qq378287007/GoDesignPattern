package main

import "fmt"

//定义电脑接口
type AbstractComputer interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

type Computer struct {
	color string
	size  int
}

func (s *Computer) SetColor(color string) {
	s.color = color
}

func (s *Computer) GetColor() string {
	return s.color
}

func (s *Computer) SetSize(size int) {
	s.size = size
}

func (s *Computer) GetSize() int {
	return s.size
}

//联想电脑
type LenovoComputer struct {
	Computer
}

//小米电脑
type XiaomiComputer struct {
	Computer
}

//定义手机接口
type AbstractPhone interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

type Phone struct {
	color string
	size  int
}

func (s *Phone) SetColor(color string) {
	s.color = color
}

func (s *Phone) GetColor() string {
	return s.color
}

func (s *Phone) SetSize(size int) {
	s.size = size
}

func (s *Phone) GetSize() int {
	return s.size
}

//联想手机
type LenovoPhone struct {
	Phone
}

//小米手机
type XiaomiPhone struct {
	Phone
}

//电子产品工厂
type InterfaceElectronicFactory interface {
	MakePhone() AbstractPhone
	MakeComputer() AbstractComputer
}

//小米品牌工厂
type XiaomiFactory struct {
}

//生产手机
func (a *XiaomiFactory) MakePhone() AbstractPhone {
	return &XiaomiPhone{
		Phone: Phone{
			color: "White",
			size:  5,
		},
	}
}

//生产电脑
func (a *XiaomiFactory) MakeComputer() AbstractComputer {
	return &XiaomiComputer{
		Computer: Computer{
			color: "Black",
			size:  14,
		},
	}
}

//联想品牌工厂
type LenovoFactory struct {
}

//生产手机
func (n *LenovoFactory) MakePhone() AbstractPhone {
	return &LenovoPhone{
		Phone: Phone{
			color: "Black",
			size:  5,
		},
	}
}

//生产电脑
func (n *LenovoFactory) MakeComputer() AbstractComputer {
	return &LenovoComputer{
		Computer: Computer{
			color: "White",
			size:  14,
		},
	}
}

//获取电子产品工厂对象
func GetElectronicFactory(brand string) (InterfaceElectronicFactory, error) {
	if brand == "Xiaomi" {
		return &XiaomiFactory{}, nil
	}

	if brand == "Lenovo" {
		return &LenovoFactory{}, nil
	}

	return nil, fmt.Errorf("%s", "error brand type")
}

func printPhoneDetails(s AbstractPhone) {
	fmt.Printf("Color: %s\n", s.GetColor())
	fmt.Printf("Size: %d inch\n", s.GetSize())
}

func printComputerDetails(s AbstractComputer) {
	fmt.Printf("Color: %s\n", s.GetColor())
	fmt.Printf("Size: %d inch\n", s.GetSize())
}

func main() {
	//声明联想工厂
	lenovoFactory, _ := GetElectronicFactory("Lenovo")

	//联想工厂生产联想手机
	lenovoPhone := lenovoFactory.MakePhone()
	printPhoneDetails(lenovoPhone)
	fmt.Println()

	//联想电脑生产联想电脑
	lenovoComputer := lenovoFactory.MakeComputer()
	printComputerDetails(lenovoComputer)
	fmt.Println()

	//声明小米工厂
	xiaomiFactory, _ := GetElectronicFactory("Xiaomi")

	//小米工厂生产小米手机
	xiaomiPhone := xiaomiFactory.MakePhone()
	printPhoneDetails(xiaomiPhone)
	fmt.Println()

	//小米电脑生产小米电脑
	xiaomiComputer := xiaomiFactory.MakeComputer()
	printComputerDetails(xiaomiComputer)
	fmt.Println()
}
