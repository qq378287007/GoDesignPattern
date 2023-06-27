package main

import "fmt"

type AlgorithmType interface {
	Delete(c *Cache)
}

//FIFO算法类型
type Fifo struct {
}

//删除缓存
func (l *Fifo) Delete(c *Cache) {
	fmt.Println("Deleting by fifo strategy")
}

//LFU算法类型
type Lfu struct {
}

//删除缓存
func (l *Lfu) Delete(c *Cache) {
	fmt.Println("Deleting by lfu strategy")
}

//LRU算法类型
type Lru struct {
}

//删除缓存
func (l *Lru) Delete(c *Cache) {
	fmt.Println("Deleting by lru strategy")
}

type Cache struct {
	storage       map[string]string
	AlgorithmType AlgorithmType
	capacity      int
	maxCapacity   int
}

func InitCache(e AlgorithmType) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:       storage,
		AlgorithmType: e,
		capacity:      0,
		maxCapacity:   2,
	}
}

func (c *Cache) SetAlgorithmType(e AlgorithmType) {
	c.AlgorithmType = e
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.Delete()
	}
	c.capacity++
	c.storage[key] = value
}

/*
func (c *Cache) Get(key string) {
	delete(c.storage, key)
}
*/

func (c *Cache) Delete() {
	c.AlgorithmType.Delete(c)
	c.capacity--
}

func main() {
	//声明Lfu对象
	lfu := &Lfu{}
	//初始化缓存对象
	cache := InitCache(lfu)
	//添加缓存
	cache.Add("one", "1")
	cache.Add("two", "2")
	cache.Add("three", "3")

	//声明Lru对象
	lru := &Lru{}
	//设置lru算法类型
	cache.SetAlgorithmType(lru)
	//添加缓存
	cache.Add("four", "4")

	//声明Fifo对象
	fifo := &Fifo{}
	//设置Fifo算法类型
	cache.SetAlgorithmType(fifo)
	//添加缓存
	cache.Add("five", "5")
}
