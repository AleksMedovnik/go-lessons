package goroutines

import (
	"fmt"
)

func Select(){
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/* 
func Training() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go f1(ch1)
	go f2(ch2)

	for i := 0; i < 2; i++ {
		// здесь должен быть твой код:)
		// если мы получим значение из ch1, вывести сообщение из ch1
		// если мы получим значение из ch2, вывести сообщение из ch2
	}
}

func f1(ch chan string) {
	ch <- "Hello from F1"
}

func f2(ch chan string) {
	ch <- "Hello from F2"
} 
*/


/* select {
case mes1 := <-ch1:
	fmt.Println(mes1)
case mes2 := <-ch2:
	fmt.Println(mes2)
} */