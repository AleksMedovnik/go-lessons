package lessons

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func Variables() {
	x := 1
	fmt.Println(x)

	var a int8
	a = 1
	fmt.Println(a)

	var b int64 = 1000
	fmt.Println(b)

	const message = "Hello"
	fmt.Println(message)

	const TOKEN = "jdieoposfdfk1355ejiefei"

	// formatted string
	message2 := fmt.Sprintf("Your name: %s, age: %d", "Alex", 21)
	fmt.Println(message2)

	// check data type
	fmt.Println(reflect.TypeOf(message)) // string
	fmt.Println(reflect.TypeOf(x))       // int

	// convert int to int
	var c int8 = 100
	d := int64(c)
	fmt.Println(reflect.TypeOf(d)) // int64

	// convert int to float
	var e int64 = 57
	var g float64 = float64(e)
	fmt.Println(reflect.TypeOf(g))

	// convert string to int
	strAge := "20"
	age, err := strconv.Atoi(strAge)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(age)) // int

	// convert string to float
	strFloat := "1.23"
	f, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(f)) // float64

	// convert int to string
	intAge := 20
	strAge = strconv.Itoa(intAge) // string

	// convert float to string
	floatNumb := 15.35
	strNumb := fmt.Sprint(floatNumb)
	fmt.Println(reflect.TypeOf(strNumb)) // string
}


// 1. Создать новый модуль с именем go_rest_api
// 2. В нем создать пакет  с функцией main.
// 3. Запросить пользователя ввести имя и возраст
// 4. Вывести сообщение: "Your name: <username>, your age: <user age>!"
// 5. Вывести сообщение: "Next year you will turn: <user age>!"
// 6. Проверить типы данных имени и возраста пользователя.
// 7. Создать переменную z для записи числа в диапазоне от -128 до 127, записать в нее число 125 и вывести в терминал.
// 8. Создать переменную u для записи числа в диапазоне от 0 до 65535, записать в нее число 125 и вывести в терминал.