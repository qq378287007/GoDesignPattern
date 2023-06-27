package main

import "fmt"

//电脑接口
type Computer interface {
	ConvertToUSB()
}

//Mac系统
type Mac struct {
}

//插入接口
func (m *Mac) ConvertToUSB() {
	fmt.Println("Lightning类型接口已插入Mac电脑")
}

//Windows操作系统
type Windows struct {
}

//插入USB接口到Windows电脑
func (w *Windows) InsertIntoUSB() {
	fmt.Println("USB接口已插入Windows电脑")
}

//Windows系统适配器
type Adapter struct {
	WindowsMachine *Windows
}

func (w *Adapter) ConvertToUSB() {
	fmt.Println("适配器将Lightning类型信号转换为USB")
	w.WindowsMachine.InsertIntoUSB()
}

//客户端
type Client struct {
}

//将Lightning类型接口插入电脑
func (c *Client) InsertIntoComputer(com Computer) {
	fmt.Println("客户端将Lightning类型接口插入计算机")
	com.ConvertToUSB()
}

func main() {
	//创建客户端
	Client := &Client{}

	//客户端插入Lightning类型连接器到Mac电脑
	Mac := &Mac{}
	Client.InsertIntoComputer(Mac)
	fmt.Println("")

	//客户端插入Lightning类型连接器到Windows适配器
	WindowsAdapter := &Windows{}
	WindowsAdapterAdapter := &Adapter{WindowsMachine: WindowsAdapter}

	Client.InsertIntoComputer(WindowsAdapterAdapter)
}
