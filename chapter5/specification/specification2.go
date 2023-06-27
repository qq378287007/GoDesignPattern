package main

import "fmt"

// 发票数据对象
type Invoice struct {
	Day    int
	Notice int
	IsSent bool
}

// 数据规格接口
type Specification interface {
	IsSatisfiedBy(Invoice) bool
	And(Specification) Specification
	Or(Specification) Specification
	Not() Specification
	Relate(Specification)
}

// 组合规格
type CompositeSpecification struct {
	Specification
}

// 检查规格
func (cs *CompositeSpecification) IsSatisfiedBy(in Invoice) bool {
	return false
}

// 规格与操作
func (cs *CompositeSpecification) And(spec Specification) Specification {
	a := &AndSpecification{
		cs.Specification, spec,
	}
	a.Relate(a)
	return a
}

// 规格或操作
func (cs *CompositeSpecification) Or(spec Specification) Specification {
	a := &OrSpecification{
		cs.Specification, spec,
	}
	a.Relate(a)
	return a
}

// 规格非操作
func (cs *CompositeSpecification) Not() Specification {
	a := &NotSpecification{
		cs.Specification,
	}
	a.Relate(a)
	return a
}

// 与规格有关
func (cs *CompositeSpecification) Relate(spec Specification) {
	cs.Specification = spec
}

// 与规格
type AndSpecification struct {
	Specification
	compare Specification
}

// 检查规格
func (as *AndSpecification) IsSatisfiedBy(in Invoice) bool {
	return as.Specification.IsSatisfiedBy(in) && as.compare.IsSatisfiedBy(in)
}

// 或规格
type OrSpecification struct {
	Specification
	compare Specification
}

// 检查规格
func (os *OrSpecification) IsSatisfiedBy(in Invoice) bool {
	return os.Specification.IsSatisfiedBy(in) || os.compare.IsSatisfiedBy(in)
}

// 非规格
type NotSpecification struct {
	Specification
}

// 检查规格
func (ns *NotSpecification) IsSatisfiedBy(in Invoice) bool {
	return ns.Specification.IsSatisfiedBy(in)
}

// 数据到期规格
type OverDueSpecification struct {
	Specification
}

// 检查规格
func (os *OverDueSpecification) IsSatisfiedBy(in Invoice) bool {
	return in.Day >= 30
}

// 创建数据到期规格
func NewOverDueSpecification() Specification {
	a := &OverDueSpecification{&CompositeSpecification{}}
	a.Relate(a)
	return a
}

// 通知发送规格
type NoticeSentSpecification struct {
	Specification
}

// 检查规格
func (ns *NoticeSentSpecification) IsSatisfiedBy(in Invoice) bool {
	return in.Notice >= 3
}

// 创建通知发送规格
func NewNoticeSentSpecification() Specification {
	a := &NoticeSentSpecification{&CompositeSpecification{}}
	a.Relate(a)
	return a
}

// 是否收到发票通知规格
type InCollectionSpecification struct {
	Specification
}

// 检查规格
func (ics *InCollectionSpecification) IsSatisfiedBy(in Invoice) bool {
	return !in.IsSent
}

// 创建是否收到发票通知规格
func NewInCollectionSpecification() Specification {
	a := &InCollectionSpecification{&CompositeSpecification{}}
	a.Relate(a)
	return a
}
func main() {
	//声明发票数据到期规格
	overDue := NewOverDueSpecification()
	//声明发票通知发送规格
	noticeSent := NewNoticeSentSpecification()
	//收款机构是否收到发票通知
	inCollection := NewInCollectionSpecification()

	sendToCollection := overDue.And(noticeSent).And(inCollection.Not())

	object := Invoice{
		Day:    32,    // >= 30
		Notice: 6,     // >= 3
		IsSent: false, // false
	}

	// 检查规格
	result := sendToCollection.IsSatisfiedBy(object)
	fmt.Println(result)
}
