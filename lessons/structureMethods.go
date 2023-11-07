package lessons

import "fmt"

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
