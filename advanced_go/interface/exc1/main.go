package main

import "fmt"

type Flyer interface {
	bird() string
	walk() string
}

type bird struct {
	Name string
}

func (b *bird) bird() string {
	return b.Name + " Flying"
}

func (b *bird) walk() string {
	return b.Name + " Walking"
}

func main() {
	b := bird{"Eagle"}

	fmt.Println(b.bird(), b.walk())
	fmt.Printf("%T\n",b)
}
