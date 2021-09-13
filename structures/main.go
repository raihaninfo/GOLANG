package main

import "fmt"

type Car struct {
	name  string
	price int
	model string
}

func (c Car) carRun() {
	fmt.Println(c.name, "is running")

}

func (c Car) getName() string {
	return c.name
}

func main() {
	c := Car{
		name:  "Tata",
		price: 25252,
		model: "AB525",
	}

	fmt.Println(c)
	c.carRun()
	fmt.Println(c.getName())

	b := struct {
		name  string
		price int
	}{
		name:  "hero",
		price: 343,
	}

	fmt.Println(b)

}
