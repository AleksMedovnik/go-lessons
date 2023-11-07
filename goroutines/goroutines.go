package goroutines

import (
	
	"fmt"
)

func Goroutines() {
	for i := 0; i < 7; i++ {
		go factorialNumber(i)
	}
	fmt.Scanln()
	fmt.Println("The End")
 

}

func factorialNumber(n int) {
	if n < 1 {
		fmt.Println("Invalid input number")
	} else if n == 1 || n == 2 {
		fmt.Println(n)
	} else {
		result := 1
		for i := 3; i <= n; i++ {
			result *= i
		}
		fmt.Println(result)
	}
}


