package main

// go mod init github.com/AleksMedovnik/go-lessons

import (
	"github.com/AleksMedovnik/go-lessons/train"
)

func main() {
	// fmt.Print("Hello", "\n")
	// fmt.Print("World", "\n")

	// fmt.Println("Hello")
	// fmt.Println("World")

	// fmt.Printf("Hello %s \n", "World!")
	// fmt.Printf("Hello %d \n", 1)
	// fmt.Printf("Hello, %f \n", 1.5)
	// fmt.Printf("Hello, %t \n", true)

	// userName := "Alex"
	// s := fmt.Sprintf("Hello, %s!", userName)
	// fmt.Println(s)

	// var name string
	// var age int
	// fmt.Print("Введите имя: ")
	// fmt.Fscan(os.Stdin, &name)

	// fmt.Print("Введите возраст: ")
	// fmt.Fscan(os.Stdin, &age)

	// fmt.Printf("В следующем году тебе исполняется: %d", age+1)

	train.Training()
	// goroutines.Mutexs()

}

