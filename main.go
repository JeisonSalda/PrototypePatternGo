package main

import (
	"fmt"

	"example.com/oslinux/model"
)

func main() {
	file1 := &model.File{Name: "File1"}
	file2 := &model.File{Name: "File2"}
	file3 := &model.File{Name: "File3"}

	folder1 := &model.Folder{
		Files: []model.INode{file1},
		Name:  "Folder1",
	}

	folder2 := &model.Folder{
		Files: []model.INode{folder1, file2, file3},
		Name:  "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("-")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("-")
}
