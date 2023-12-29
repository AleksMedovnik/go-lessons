package lessons

import (
	"fmt"
	"os"
)

func Maps() {
	user := map[string]string{
		"name":    "Alex",
		"country": "Russia",
		"city":    "Moscow",
	}
	fmt.Println(user["name"]) // Alex

	if val, ok := user["city"]; ok {
		fmt.Println(ok)
		fmt.Println(val)
	}

	for key, value := range user {
		fmt.Printf("%s: %s\n", key, value)
	}

	customer := make(map[string]string)
	customer["name"] = "Alex"
	customer["country"] = "Russia"
	delete(customer, "country")
	fmt.Println(customer)

	customer["city"] = "New York"
	fmt.Println(customer["city"]) // New York
	changeUser(customer, "Moscow")
	fmt.Println(customer["city"]) // Moscow

	// ######################
	phoneBook := map[string]string{
		"Vasya": "+7 959 331 25 00",
		"Pete": "+7 959 331 25 01",
		"Masha": "+7 959 331 25 02",
	}

	showPhoneNumber(phoneBook)
	// Input: Vasya
	// Output: +7 959 331 25 00

	showPhoneNumber(phoneBook)
	// Input: Alex
	// Output: There is no user with that name in the book

}

func changeUser(user map[string]string, city string) {
	user["city"] = city
}

// #################################

// Дан map:
/* phoneBook := map[string]string{
	"Vasya": "+7 959 331 25 00",
	"Pete": "+7 959 331 25 01",
	"Masha": "+7 959 331 25 02",
} */

// Пользователь вводит имя в терминал и должен получить номер телефона
// Example:
// showPhoneNumber(phoneBook)
// Input: Vasya
// Output: +7 959 331 25 00

// Если пользователь отсутствует в списке, написать соответствующее сообщение.
// Example:
// showPhoneNumber(phoneBook)
// Input: Alex
// Output: There is no user with that name in the book

func showPhoneNumber(phoneBook map[string]string) {
	var username string
    fmt.Print("Enter username: ")
	fmt.Fscan(os.Stdin, &username)
	if value, ok := phoneBook[username]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("There is no user with that name in the book")
	}
	
}
