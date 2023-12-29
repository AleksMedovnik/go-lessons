package train

import "fmt"

func Training() {
	intCh := make(chan int)

	go factorial(intCh)

	for i := 1; i <= 3; i++ {
		fmt.Println(<-intCh)
	}
}

func factorial(ch chan int) {
	result, count := 1, 1
	for {
		result *= count
		count++
		ch <- result 
	}
}


/* func fibonacci(ch chan int) {
	a, b := 0, 1
	for {
		a, b = b, a + b
		ch <- a
	}
} */