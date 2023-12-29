package goroutines

import (
	"fmt"
)

func Goroutines() {
	for i := 0; i < 7; i++ {
		go sum(i, i)
	}
	fmt.Scanln()
	fmt.Println("The End")
}

func sum(a, b int) {
	fmt.Println(a + b)
}


