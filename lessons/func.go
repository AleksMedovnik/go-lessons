package lessons

import (
	"fmt"
	"os"
	"errors"
)

func Funcions() {
	x := 0
	changeValue(&x)
	fmt.Println(x) // 1

	f := func() {
		x++
	}
	f()
	fmt.Println(x) // 2

	res, err := Sum(150, 150)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(res) // 3

}

func loops() {
	var x int16 = 3
	var y string
	switch x {
	case 1:
		y = "One"
	case 2:
		y = "Two"
	default:
		y = fmt.Sprint(x)
	}
	fmt.Println(y)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	validPassword := "123abc"
	var userPassword string

	for true {
		fmt.Print("Enter your password: ")
		fmt.Fscan(os.Stdin, &userPassword)

		if userPassword == validPassword {
			fmt.Println("Welcome!")
			break
		}
	}
}

func Sum(a, b byte) (byte, error) {
	return a + b, nil
}

// неопределенное количество параметров
func sum_args(params ...int) int {
	var result int = 0
	for _, param := range params {
		result += param
	}
	return result
}

// Возвращение нескольких значений
func get_size(width, height int) (int, int) {
	return width * 2, height * 2
}

func changeValue(x *int) {
	*x++
}

func divide(x, y float64) float64 {
	if y == 0 {
		panic("Division by zero!") // генерирует ошибку
	}
	return x / y
}

func sayHello(userName string, userAge int, isDeveloper bool) string {
	return fmt.Sprintf("Hello, %s! Your age: %d and you are %t", userName, userAge, isDeveloper)
}

func factorial(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Invalid input number")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}


/*  Написать функцию sumAll(a), которая работает следующим образом:
sumAll(1)(2) // 1 3 (Вывод в терминал по очереди)
*/

/*  Написать функцию showParams, которая принимает неопределенное количество параметров 
с именами пользователей и возвращает строку с именами при условии, что имя длинее 2-х символов.
Example: 
names := showParams("Hi", "Alex", "Maxim", "X", "Masha")
fmt.Println(names) // Alex Maxim Masha
*/

func showParams(names ...string) string {
	var result string = ""
	for _, param := range names {
		if len(param) > 2 {
			result += param + " "
		}
	}
	return result
}

func sumAll(a int) func(b int) {
	fmt.Println(a)
	return func(b int) {
		a += b
		fmt.Println(a)
	}
}