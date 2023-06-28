package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Iterator interface {
	HasNext() bool
	GetNext() *User
}

type Collection interface {
	CreateIterator() Iterator
}

type UserIterator struct {
	Index int
	Users []*User
}

func (u *UserIterator) HasNext() bool {
	if u.Index < len(u.Users) {
		return true
	}
	return false

}

func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.Users[u.Index]
		u.Index++
		return user
	}
	return nil
}

type UserCollection struct {
	Users []*User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{Users: u.Users}
}

func main() {
	//声明用户对象User1
	user1 := &User{
		Name: "Jack",
		Age:  30,
	}
	//声明用户对象User2
	user2 := &User{
		Name: "Barry",
		Age:  20,
	}

	//声明具体集合对象
	userCollection := &UserCollection{Users: []*User{user1, user2}}

	//声明具体迭代器对象
	iterator := userCollection.CreateIterator()

	//执行具体方法
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
