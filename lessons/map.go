package lessons

import (
	"fmt"
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
}

func changeUser(user map[string]string, city string) {
	user["city"] = city
}
