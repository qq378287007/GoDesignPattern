package main

import "fmt"

type Component interface {
	Search(string)
}

type File struct {
	Name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("在文件 %s 中递归搜索关键 %s \n", f.Name, keyword)
}

func (f *File) GetName() string {
	return f.Name
}

type Folder struct {
	Components []Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("在文件夹 %s 中递归搜索关键 %s \n", f.Name, keyword)
	for _, composite := range f.Components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.Components = append(f.Components, c)
}

func main() {
	File1 := &File{Name: "File1"}
	File2 := &File{Name: "File2"}
	File3 := &File{Name: "File3"}

	Folder1 := &Folder{Name: "Folder1"}
	Folder1.Add(File1)

	Folder2 := &Folder{Name: "Folder2"}
	Folder2.Add(File2)
	Folder2.Add(File3)
	Folder2.Add(Folder1)

	Folder2.Search("keyword")
}
