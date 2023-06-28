package main

import "fmt"

type Patient struct {
	Name              string
	RegistrationDone  bool
	ClinicCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

type department interface {
	Execute(*Patient)
	SetNext(department)
}

type Cashier struct {
	next department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("支付完成")
	}
	fmt.Println("收银员从病人那里收钱")
}

func (c *Cashier) SetNext(next department) {
	c.next = next
}

type Clinic struct {
	next department
}

func (d *Clinic) Execute(p *Patient) {
	if p.ClinicCheckUpDone {
		fmt.Println("医生已经检查过了")
		d.next.Execute(p)
		return
	}
	fmt.Println("医生正在检查病人")
	p.ClinicCheckUpDone = true
	d.next.Execute(p)
}

func (d *Clinic) SetNext(next department) {
	d.next = next
}

type Drugstore struct {
	next department
}

func (m *Drugstore) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("药品已经给病人")
		m.next.Execute(p)
		return
	}
	fmt.Println("正在给病人用药")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Drugstore) SetNext(next department) {
	m.next = next
}

type Reception struct {
	next department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("已完成患者登记")
		r.next.Execute(p)
		return
	}
	fmt.Println("正在接待登记病人")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next department) {
	r.next = next
}

func main() {

	cashier := &Cashier{}

	//设置下一个医务部门
	medical := &Drugstore{}
	medical.SetNext(cashier)

	//设置下一个医务部门
	doctor := &Clinic{}
	doctor.SetNext(medical)

	//设置下一个医务部门
	reception := &Reception{}
	reception.SetNext(doctor)

	patient := &Patient{Name: "Jack"}
	//设置病人
	reception.Execute(patient)
}
