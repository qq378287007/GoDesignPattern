package main

import (
	"fmt"
	"log"
)

type State interface {
	AddProduct(int) error
	RequestProduct() error
	InsertMoney(money int) error
	DispenseProduct() error
}

type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (i *HasMoneyState) RequestProduct() error {
	return fmt.Errorf("Product dispense in progress")
}

func (i *HasMoneyState) AddProduct(count int) error {
	return fmt.Errorf("Product dispense in progress")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Product out of stock")
}
func (i *HasMoneyState) DispenseProduct() error {
	fmt.Println("Dispensing Product")
	i.VendingMachine.productCount = i.VendingMachine.productCount - 1
	if i.VendingMachine.productCount == 0 {
		i.VendingMachine.SetState(i.VendingMachine.noProduct)
	} else {
		i.VendingMachine.SetState(i.VendingMachine.hasProduct)
	}
	return nil
}

type HasProductState struct {
	VendingMachine *VendingMachine
}

func (i *HasProductState) RequestProduct() error {
	if i.VendingMachine.productCount == 0 {
		i.VendingMachine.SetState(i.VendingMachine.noProduct)
		return fmt.Errorf("No product present")
	}
	fmt.Printf("Product requestd\n")
	i.VendingMachine.SetState(i.VendingMachine.productRequested)
	return nil
}

func (i *HasProductState) AddProduct(count int) error {
	fmt.Printf("%d products added\n", count)
	i.VendingMachine.IncrementProductCount(count)
	return nil
}

func (i *HasProductState) InsertMoney(money int) error {
	return fmt.Errorf("Please select product first")
}
func (i *HasProductState) DispenseProduct() error {
	return fmt.Errorf("Please select product first")
}

type ProductRequestedState struct {
	VendingMachine *VendingMachine
}

func (i *ProductRequestedState) RequestProduct() error {
	return fmt.Errorf("Product already requested")
}

func (i *ProductRequestedState) AddProduct(count int) error {
	return fmt.Errorf("Product Dispense in progress")
}

func (i *ProductRequestedState) InsertMoney(money int) error {
	if money < i.VendingMachine.productPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", i.VendingMachine.productPrice)
	}
	fmt.Println("Money entered is ok")
	i.VendingMachine.SetState(i.VendingMachine.hasMoney)
	return nil
}
func (i *ProductRequestedState) DispenseProduct() error {
	return fmt.Errorf("Please insert money first")
}

type NoProductState struct {
	VendingMachine *VendingMachine
}

func (i *NoProductState) RequestProduct() error {
	return fmt.Errorf("Product out of stock")
}

func (i *NoProductState) AddProduct(count int) error {
	i.VendingMachine.IncrementProductCount(count)
	i.VendingMachine.SetState(i.VendingMachine.hasProduct)
	return nil
}

func (i *NoProductState) InsertMoney(money int) error {
	return fmt.Errorf("Product out of stock")
}
func (i *NoProductState) DispenseProduct() error {
	return fmt.Errorf("Product out of stock")
}

// 自动售货机类
type VendingMachine struct {
	hasProduct       State
	productRequested State
	hasMoney         State
	noProduct        State

	currentState State

	productCount int
	productPrice int
}

func NewVendingMachine(productCount, productPrice int) *VendingMachine {
	v := &VendingMachine{
		productCount: productCount,
		productPrice: productPrice,
	}
	HasProductState := &HasProductState{VendingMachine: v}
	ProductRequestedState := &ProductRequestedState{VendingMachine: v}
	HasMoneyState := &HasMoneyState{VendingMachine: v}
	NoProductState := &NoProductState{VendingMachine: v}

	v.SetState(HasProductState)
	v.hasProduct = HasProductState
	v.productRequested = ProductRequestedState
	v.hasMoney = HasMoneyState
	v.noProduct = NoProductState
	return v
}

func (v *VendingMachine) RequestProduct() error {
	return v.currentState.RequestProduct()
}

func (v *VendingMachine) AddProduct(count int) error {
	return v.currentState.AddProduct(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseProduct() error {
	return v.currentState.DispenseProduct()
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) IncrementProductCount(count int) {
	fmt.Printf("Adding %d products\n", count)
	v.productCount = v.productCount + count
}

func main() {
	//声明自动售货机
	VendingMachine := NewVendingMachine(1, 10)

	//请求商品
	err := VendingMachine.RequestProduct()
	if err != nil {
		log.Fatalf(err.Error())
	}

	//向自动售货机放入10元钱
	err = VendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	//向自动售货机分配商品
	err = VendingMachine.DispenseProduct()
	if err != nil {
		log.Fatalf(err.Error())
	}

	//向自动售货机添加2个商品
	err = VendingMachine.AddProduct(2)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
