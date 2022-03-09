package model

import "fmt"

type File struct {
	Name string
}

func (file *File) Print(indentation string) {
	fmt.Println(indentation + file.Name)
}

func (file *File) Clone() INode {
	return &File{Name: file.Name + "_clone"}
}
