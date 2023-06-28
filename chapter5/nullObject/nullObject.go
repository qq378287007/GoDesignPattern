package main

import "fmt"

//部门接口
type department interface {
	getNumberOfProfessors() int
	getName() string
}

//计算机科学系
type computerscience struct {
	numberOfProfessors int
}

//获取计算机科学系教授数量
func (c *computerscience) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

//获取部门名字
func (c *computerscience) getName() string {
	return "computerscience"
}

//机械系
type mechanical struct {
	numberOfProfessors int
}

//获取机械系教授数量
func (c *mechanical) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

//获取部门名字
func (c *mechanical) getName() string {
	return "mechanical"
}

//空对象
type nullDepartment struct {
	numberOfProfessors int
}

//获取空对象教授数量
func (c *nullDepartment) getNumberOfProfessors() int {
	return 0
}

//获取空对象名字
func (c *nullDepartment) getName() string {
	return ""
}

//学院
type college struct {
	departments []department
}

//添加部门
func (c *college) addDepartment(departmentName string, numOfProfessors int) {
	if departmentName == "computerscience" {
		computerScienceDepartment := &computerscience{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, computerScienceDepartment)
	}
	if departmentName == "mechanical" {
		mechanicalDepartment := &mechanical{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, mechanicalDepartment)
	}
}

//获取部门
func (c *college) getDepartment(departmentName string) department {
	for _, department := range c.departments {
		if department.getName() == departmentName {
			return department
		}
	}
	//如果部门不存在，则返回空部门
	return &nullDepartment{}
}

//创建学院1
func createCollege1() *college {
	college := &college{}
	college.addDepartment("computerscience", 0)
	college.addDepartment("mechanical", 2)
	return college
}

//创建学院2
func createCollege2() *college {
	college := &college{}
	college.addDepartment("computerscience", 3)
	return college
}

func main() {
	college1 := createCollege1()
	totalProfessors := 0
	departmentArray := []string{"computerscience", "mechanical", "chinese", "computer"}
	for _, deparmentName := range departmentArray {
		d := college1.getDepartment(deparmentName)
		totalProfessors += d.getNumberOfProfessors()
	}
	fmt.Printf("学院1的教授数量为： %d\n", totalProfessors)

	college2 := createCollege2()
	totalProfessors = 0
	for _, deparmentName := range departmentArray {
		d := college2.getDepartment(deparmentName)
		totalProfessors += d.getNumberOfProfessors()
	}
	fmt.Printf("学院2的教授数量为： %d\n", totalProfessors)
}
