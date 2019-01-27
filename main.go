package main

import "fmt"

// Directory ...
type Directory interface {
	Print() string
}

// Location ...
type Location struct {
	Type string
}

// Salary ...
type Salary struct {
}

// Category ...
type Category struct {
}

// Subcategory ...
type Subcategory struct {
}

// Keyword ...
type Keyword struct {
}

// NewDirectory ...
func NewDirectory(text string) Directory {

	switch text {
	case "location":
		return Location{"loc"}
	case "salary":
		return Salary{}
	default:
		return Keyword{}
	}
}

// Print Location
func (l Location) Print() string {
	return "Soy una Locación; Type: " + l.Type
}

// Print Salary
func (s Salary) Print() string {
	return "Soy un Salario"
}

// Print Keyword
func (k Keyword) Print() string {
	return "Soy un Keyword por Default"
}

func main() {
	queretaro := NewDirectory("location")
	benito := NewDirectory("salary")
	indefinido := NewDirectory("hola")

	fmt.Println(queretaro.Print())
	fmt.Println(benito.Print())
	fmt.Println(indefinido.Print())
}
