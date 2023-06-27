package main

import (
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"
)

const getResMaxTime = 3 * time.Second

var (
	ErrPoolNotExist  = errors.New("pool not exist")
	ErrGetResTimeout = errors.New("get resource time out")
)

// 资源类
type Resource struct {
	reusable int
}

// 初始化资源对象
// 模拟缓慢的资源访问，例如，TCP 连接等
func NewResource(id int) *Resource {
	time.Sleep(500 * time.Millisecond)
	return &Resource{reusable: id}
}

// 模拟资源耗时
func (r *Resource) Do(workId int) {
	time.Sleep(time.Duration(rand.Intn(5)) * 100 * time.Millisecond)
	log.Printf("using resource #%d finished work %d finish\n", r.reusable, workId)
}

// 对象池
type Pool chan *Resource

// 并发创建资源对象，节省资源对象初始化时间
func New(size int) Pool {
	p := make(Pool, size)
	wg := new(sync.WaitGroup)
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(reusable int) {
			p <- NewResource(reusable)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return p
}

// 从获取对象池获取对象
func (p Pool) GetResource() (r *Resource, err error) {
	select {
	case r := <-p:
		return r, nil
	case <-time.After(getResMaxTime):
		return nil, ErrGetResTimeout
	}
}

// 将资源返回到资源池
func (p Pool) GiveBackResource(r *Resource) error {
	if p == nil {
		return ErrPoolNotExist
	}
	p <- r
	return nil
}

func main() {
	// 初始化一个包含五个资源的资源池
	// 可以调整为 1 或 10 以查看差异
	size := 5
	p := New(size)

	// 调用资源池
	doWork := func(workId int, wg *sync.WaitGroup) {
		defer wg.Done()
		// 从资源池中获取资源对象
		res, err := p.GetResource()
		if err != nil {
			log.Println(err)
			return
		}
		//返回的资源对象
		defer p.GiveBackResource(res)
		// 使用资源处理工作
		res.Do(workId)
	}

	// 模拟100个并发进程从资产池中获取资源对象
	num := 100
	wg := new(sync.WaitGroup)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go doWork(i, wg)
	}
	wg.Wait()
}
