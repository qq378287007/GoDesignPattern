package main

import "fmt"

//打印机接口
type Printer interface {
	PrintFile()
}

//联想打印机
type Lenovo struct {
}

//打印文件
func (p *Lenovo) PrintFile() {
	fmt.Println("Printing by a Lenovo Printer")
}

//佳能打印机
type Canon struct {
}

//打印文件
func (p *Canon) PrintFile() {
	fmt.Println("Printing by a Canon Printer")
}

//电脑接口
type Computer interface {
	Print()
	SetPrinter(Printer)
}

//Mac系统
type Mac struct {
	Printer Printer
}

//打印
func (m *Mac) Print() {
	fmt.Println("Print request for Mac")
	m.Printer.PrintFile()
}

//设置打印机
func (m *Mac) SetPrinter(p Printer) {
	m.Printer = p
}

//Windows系统
type Windows struct {
	Printer Printer
}

//打印
func (w *Windows) Print() {
	fmt.Println("Print request for Windows")
	w.Printer.PrintFile()
}

//设置打印机
func (w *Windows) SetPrinter(p Printer) {
	w.Printer = p
}

func main() {
	//联想打印机
	var lenovoPrinter Printer = &Lenovo{}

	//佳能打印机
	var canonPrinter Printer = &Canon{}

	//Mac打印
	var macComputer Computer = &Mac{}

	//Mac电脑用SetPrinter()方法设置联想打印机
	macComputer.SetPrinter(lenovoPrinter)
	macComputer.Print()
	fmt.Println()

	//Mac电脑用SetPrinter()方法设置佳能打印机
	macComputer.SetPrinter(canonPrinter)
	macComputer.Print()
	fmt.Println()

	var winComputer Computer = &Windows{}
	//Windows电脑用SetPrinter()方法设置联想打印机
	winComputer.SetPrinter(lenovoPrinter)
	winComputer.Print()
	fmt.Println()

	///Windows电脑用SetPrinter()方法设置佳能打印机
	winComputer.SetPrinter(canonPrinter)
	winComputer.Print()
	fmt.Println()
}
