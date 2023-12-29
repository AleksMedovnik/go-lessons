package lessons

import (
	"fmt"
)

type mile int

type user struct {
	name string
	age  int
}

type address struct {
	country string
	city    string
}

type person struct {
	name string
	age  int
	address
}

// структура односвязного списка
type node struct {
	value int
	next  *node
}

func CustomType() {
	var distance mile = 150
	fmt.Println(distance) // 150
	distance++
	fmt.Println(distance) // 151
	// distance2 := 20
	// fmt.Println(distanceCompare(distance, distance2)) // Error

	var distance2 mile = 20
	fmt.Println(distanceCompare(distance, distance2)) // 131
	fmt.Println(distanceCompare(distance2, distance)) // 131
}

func distanceCompare(distance1, distance2 mile) mile {
	if distance1 > distance2 {
		return distance1 - distance2
	} else {
		return distance2 - distance1
	}
}

func Structure() {
	alex := user{"Alex", 21}
	fmt.Println(alex.name) // Alex

	alexPointer := &alex
	alexPointer.age = 30
	fmt.Println(alex.age)        // 30
	fmt.Println(alexPointer.age) // 30

	var bob *user = new(user)
	fmt.Println(bob.name) // ""
	bob.name = "Bob"
	fmt.Println(bob.name) // Bob

	u := person{
		name: "Alex",
		age:  21,
		address: address{
			country: "Russia",
			city:    "Moscow",
		},
	}

	fmt.Println(u.address.country) // Russia
	fmt.Println(u.country)         // Russia

	first := node{value: 1}
	second := node{value: 2}
	third := node{value: 3}

	first.next = &second
	second.next = &third

	printNodeValue(&first)

	// #######################
	// productList := []*product{
	// 	product{"Iphone 13", 999, "About iphone 13"},
	// 	product{"Iphone 14", 1050, "About iphone 14"},
	// 	product{"Samsung Galaxy", 750, "Samsung Galaxy"},
	// }

	products := []*Product{
		{"Iphone 13", "About Iphone 13", 999},
		{"Iphone 14", "About Iphone 14", 1200},
		{"Samsung Galaxy", "About Samsung", 800},
	}

	fmt.Println(averPrice(products)) // 999

}

func printNodeValue(n *node) {
	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}

// ########################################
type Product struct {
	title string
	info  string
	price int
}

// написать функцию averPrice, принимающую slice со структурами
// Product и возвращающую среднюю стоимость продуктов

// ###############################
func averPrice(p []*Product) int {
	res := 0
	for _, product := range p {
		res += product.price
	}
	return res / len(p)
}
