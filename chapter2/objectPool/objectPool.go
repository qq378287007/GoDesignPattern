package main

import (
	"fmt"
	"sync"
)

// 对象池
type Pool struct {
	sync.Mutex
	Inuse     []interface{}
	Available []interface{}
	new       func() interface{}
}

// 创建一个新对象池
func NewPool(new func() interface{}) *Pool {
	return &Pool{new: new}
}

// 从池中获取要使用的新池对象。
// 如果没有可用，则获取创建1个池对象的新实例
func (p *Pool) Acquire() interface{} {
	p.Lock()
	var object interface{}
	if len(p.Available) != 0 {
		object = p.Available[0]
		p.Available = append(p.Available[:0], p.Available[1:]...)
	} else {
		object = p.new()
	}
	p.Inuse = append(p.Inuse, object)
	p.Unlock()
	return object
}

// 将对象释放回对象池
func (p *Pool) Release(object interface{}) {
	p.Lock()
	p.Available = append(p.Available, object)
	for i, v := range p.Inuse {
		if v == object {
			p.Inuse = append(p.Inuse[:i], p.Inuse[i+1:]...)
			break
		}
	}
	p.Unlock()
}

func main() {
	num := func() interface{} {
		return 10.0
	}
	pool := NewPool(num)
	object := pool.Acquire()

	fmt.Println(pool.Inuse)
	fmt.Println(pool.Available)

	pool.Release(object)
	fmt.Println(pool.Inuse)
	fmt.Println(pool.Available)
}
