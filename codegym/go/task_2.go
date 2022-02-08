package main

import (
	"fmt"
)

type Circle struct {
	radius float32
	pi     float32
}

type Person struct {
	name                       string
	weight, height, age, money int
}

// can not use 'wift Woman' directly
// will cause the size and memory layout of the struct becomes recursive
// use pointer
type Man struct {
	Person
	wife *Woman
}

type Woman struct {
	Person
	husband *Man
}

type Fahrenheit float32
type Celsius float32

func (c Circle) printCircleCircumference() {
	fmt.Println(2.0 * c.radius * c.pi)
}

// get age function
func (p Person) GetAge() int {
	return p.age
}

// custom string function like java toString()
func (m Man) String() string {
	return fmt.Sprintf("%+v wife:{%+v}", m.Person, m.wife.Person)
}

func (w Woman) String() string {
	return fmt.Sprintf("%+v husband:{%+v}", w.Person, w.husband.Person)
}

func convertCelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func convertFahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	c := Circle{radius: 5, pi: 3.14}
	c.printCircleCircumference()
	fmt.Println(convertCelsiusToFahrenheit(41))
	fmt.Println(convertFahrenheitToCelsius(105.8))

	p := new(Person)
	p.age = 2
	p.name = "d"
	fmt.Printf("%+v\n", p)

	woman := Woman{Person: Person{name: "Alice"}}
	man := Man{Person: Person{name: "John"}}
	woman.age = 18
	man.age = 25
	woman.husband = &man
	man.wife = &woman

	fmt.Printf("%+v\n", woman)
	// fmt.Println(woman.husband.name)
	fmt.Printf("%+v\n", man)
	// fmt.Println(man.wife.name)
}
