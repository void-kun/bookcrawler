package main

import "fmt"

// import (
// 	"bws/crawl"
// )

// func main() {
// 	metadataPath := "./sources.json"
// 	myCrawl := crawl.Crawl{}

// 	myCrawl.LoadSources(metadataPath)
// }
//

type AbstractInterface interface {
	SomeFunction() string
}

type Source struct {
	Name string
	Age  string
}

type ConcreteStruct1 struct {
	Source
}

func (c *ConcreteStruct1) SomeFunction() string {
	c.Name = "3333"
    fmt.Println("ConcreteStruct 1, name = ", c.Name)
    return c.Name
}

func (c *ConcreteStruct1) New(name string, sou string) *ConcreteStruct1 {
	c.Name = name
	c.Age = sou
	return c
}

type ConcreteStruct2 struct {
	Source
}

func (c ConcreteStruct2) SomeFunction() string {
	fmt.Println("ConcreteStruct 2")
    return c.Age
}

func main() {
	var array []AbstractInterface

	array = append(array, &ConcreteStruct1{Source{Name: "hello", Age: "32"}})
    array = append(array, ConcreteStruct2{Source {Name: "value", Age: "323"}})
	array = append(array, &ConcreteStruct1{})

	for _, item := range array {
		item.SomeFunction()
	}

	fmt.Println("name of struct 1, ", (array[0].SomeFunction()))
}
