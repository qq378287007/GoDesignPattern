package main

import "fmt"

//节点接口
type InterfaceNode interface {
	Print(string)
	Clone() InterfaceNode
}

//文件类型
type File struct {
	Name string
}

//打印
func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

//克隆
func (f *File) Clone() InterfaceNode {
	return &File{Name: f.Name + "_Clone"}
}

//文件夹
type Folder struct {
	Children []InterfaceNode
	Name     string
}

//打印
func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	//fmt.Println(f.Name)
	for _, i := range f.Children {
		i.Print(indentation + indentation)
		//i.Print(indentation)
	}
}

//克隆
func (f *Folder) Clone() InterfaceNode {
	CloneFolder := &Folder{Name: f.Name + "_Clone"}
	var tempChildren []InterfaceNode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	CloneFolder.Children = tempChildren
	return CloneFolder
}

func main() {
	//声明文件对象File1
	File1 := &File{Name: "File1"}
	//声明文件对象File2
	File2 := &File{Name: "File2"}
	//声明文件对象File3
	File3 := &File{Name: "File3"}

	//声明文件夹对象Folder1
	Folder1 := &Folder{
		Children: []InterfaceNode{File1},
		Name:     "文件夹Folder1",
	}

	//声明文件夹对象Folder2
	Folder2 := &Folder{
		Children: []InterfaceNode{Folder1, File2, File3},
		Name:     "文件夹Folder2",
	}
	fmt.Println("\n打印文件夹Folder2的层级:")
	Folder2.Print("  ")

	CloneFolder := Folder2.Clone()
	fmt.Println("\n打印复制文件夹Folder2的层级:")
	CloneFolder.Print("  ")
}
