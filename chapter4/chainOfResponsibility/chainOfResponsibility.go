package main

import "fmt"

// Handler 定义了一个处理程序来处理给定的 handleID
type Handler interface {
	SetNext(handler Handler)
	Handle(handleID int) int
}

//基础处理者
type BaseHandler struct {
	name     string
	next     Handler
	handleID int
}

//NewHandler 返回一个新的处理程序
func NewBaseHandler(name string, next Handler, handleID int) Handler {
	return &BaseHandler{name, next, handleID}
}

// Handle 处理给定的 handleID
func (h *BaseHandler) Handle(handleID int) int {
	if handleID < 4 {
		ch := &ConcreteHandler{}
		ch.Handle(handleID)
		fmt.Println(h.name)

		if h.next != nil {
			h.next.Handle(handleID + 1)
		}

		return handleID + 1
	}
	return 0
}

// 设置下一个处理者
func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

//具体处理者
type ConcreteHandler struct {
}

//具体处理者的处理方法
func (ch *ConcreteHandler) Handle(handleID int) {
	fmt.Println("ConcreteHandler handleID:", handleID)
}

func main() {
	barry := NewBaseHandler("Barry", nil, 1)
	shirdon := NewBaseHandler("Shirdon", barry, 2)
	jack := NewBaseHandler("Shirdon", shirdon, 3)

	res := shirdon.Handle(2)
	fmt.Println(res)

	fmt.Println()

	res1 := jack.Handle(3)
	fmt.Println(res1)
}
