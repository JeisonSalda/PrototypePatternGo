package model

type INode interface {
	Print(string)
	Clone() INode
}
