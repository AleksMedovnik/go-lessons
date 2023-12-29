package lessons

import (
	"fmt"
)

type lib []string

type usr struct {
	name string
	age  int
}

func StructureMethods() {
	library := lib{"Book1", "Book2", "Book3"}
	library.print()

	u := usr{"Masha", 20}
	u.displayInfo() // Username: Masha, age: 20

	// Методы указателей
	fmt.Printf("User age before update: %d\n", u.age) // User age before update: 20
	u.updateAge(25)
	fmt.Printf("User age after update: %d\n", u.age) // User age after update: 25

	p := product{"Iphone 14", "About iphone 14", 999}
	fmt.Println(p.showPrice()) // Price: $999
}

func (l lib) print() {
	for _, val := range l {
		fmt.Println(val)
	}
}

func (u usr) displayInfo() {
	fmt.Printf("Username: %s, age: %d\n", u.name, u.age)
}

func (u *usr) updateAge(newAge int) {
	u.age = newAge
}

// ##########################################
// 1. Определить структуру для продукта. В ней должны быть поля для названия, 
// описания и стоимости товара с целочисленным значением. 
// Добавить метод showPrice, возвращающий сообщение о стоимости товара.
// Example:
// p := product{"Iphone 14", "About iphone 14", 999}
// fmt.Println(p.showPrice()) // Price: $999

// ###########################
type product struct {
	title string
	info string
	price int
}

func (p product) showPrice() string {
	return fmt.Sprintf("Price: $%d", p.price)
}


// 2. Создать структуру mobilePhone, которая имеет те же атрибуты, что и product,
// но, также, атрибуты maker и os. 
// Создать структуру closes. Она тоже имеет атрибуты из product. 
// Но, также, maker и size.
// Попробуй правильно распределить свойства, чтобы не было дублирования кода