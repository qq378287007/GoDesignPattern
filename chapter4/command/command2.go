package main

import "fmt"

type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.Execute()
}

type Command interface {
	Execute()
}

type Device interface {
	On()
	Off()
}

type Light struct {
	isRunning bool
}

func (t *Light) On() {
	t.isRunning = true
	fmt.Println("打开灯...")
}

func (t *Light) Off() {
	t.isRunning = false
	fmt.Println("关闭灯...")
}

type OffCommand struct {
	Device Device
}

func (c *OffCommand) Execute() {
	c.Device.Off()
}

type OnCommand struct {
	Device Device
}

func (c *OnCommand) Execute() {
	c.Device.On()
}

func main() {
	//初始化具体接收者对象
	Light := &Light{}

	//发送打开命令
	onCommand := &OnCommand{
		Device: Light,
	}

	//发送关闭命令
	offCommand := &OffCommand{
		Device: Light,
	}

	//接收打开命令
	onButton := &Button{
		Command: onCommand,
	}
	//按打开命令键
	onButton.Press()

	//接收关闭命令
	offButton := &Button{
		Command: offCommand,
	}
	//按关闭命令键
	offButton.Press()
}
