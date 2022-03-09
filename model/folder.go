package model

import "fmt"

type Folder struct {
	Files []INode
	Name  string
}

func (folder *Folder) Print(indentation string) {
	fmt.Println(indentation + folder.Name)
	for _, file := range folder.Files {
		file.Print(indentation + indentation)
	}
}

func (folder *Folder) Clone() INode {
	cloneFolder := &Folder{Name: folder.Name + "_clone"}
	var tempFiles []INode
	for _, file := range folder.Files {
		copy := file.Clone()
		tempFiles = append(tempFiles, copy)
	}
	cloneFolder.Files = tempFiles
	return cloneFolder
}
