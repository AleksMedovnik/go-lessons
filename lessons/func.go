package lessons

import (
	"errors"
	"fmt"
	"os"
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

	var a float64 = 15.0
	var b float64 = 0
	res, err := Divide(a, b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res) // 3

	fact, err := Factorial(0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(fact) // -1

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

func Divide(x, y float64) (float64, error) {
	if y == 0 {
		return -1, errors.New("Division by zero!")
	}
	return x / y, nil
}

func sayHello(userName string, userAge int, isDeveloper bool) string {
	return fmt.Sprintf("Hello, %s! Your age: %d and you are %t", userName, userAge, isDeveloper)
}

func Factorial(n int64) (int64, error) {
	if n < 1 {
		return -1, errors.New("Invalid input number")
	} else if n == 1 {
		return n, nil
	}
	var i int64
	var result int64 = 1
	for i = 2; i <= n; i++ {
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
