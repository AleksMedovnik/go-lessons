package train

import "fmt"

func Training() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go f1(ch1)
	go f2(ch2)

	for i := 0; i < 2; i++ {
		select {
		case mes1 := <-ch1:
			fmt.Println(mes1)
		case mes2 := <-ch2:
			fmt.Println(mes2)
		}
	}
}

func f1(ch chan string) {
	ch <- "Hello from F1"
}

func f2(ch chan string) {
	ch <- "Hello from F2"
}
