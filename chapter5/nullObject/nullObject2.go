package main

import (
	"errors"
	"fmt"
)

// 抽象对象
type AbstractObject interface {
	Request(str string) (string, error)
}

// 空对象
type NullObject struct {
}

func (w *NullObject) Request(str string) (string, error) {
	return "null", errors.New("not implemented yet!")
}

// 真实对象
type RealObject struct {
}

func (w *RealObject) Request(str string) (string, error) {
	return str, nil
}

func main() {
	objMap := make(map[string]bool)
	//objMap["real"] = true

	var object AbstractObject
	if objMap["real"] {
		object = &RealObject{}
	} else {
		object = &NullObject{}
	}
	fmt.Println(object.Request("real"))
}
