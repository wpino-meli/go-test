package library

import "fmt"

type library struct{}

type ILibrary interface {
	Hello()
}

func NewLibrary() ILibrary {
	return library{}
}

func (l library) Hello() {
	fmt.Println("I'am in library 1")
}
