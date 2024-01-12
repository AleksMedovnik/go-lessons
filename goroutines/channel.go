package goroutines

import (
	"fmt"
	"os"
)

func Channel() {
	// Небуферизированный канал
	intCh := make(chan int) // канал для данных типа int

	// получение данных из анонимной функции
	go func() {
		fmt.Println("Goroutine start")
		intCh <- 5 // блокировка, пока данные не будут получены функцией main
	}()

	fmt.Println(<-intCh)

	// получение данных из именованной функции
	c := make(chan int)
	go factorial(11, c)
	fmt.Println(<-c) // 39916800
	// ###############################

	// Буферизированный канал
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	fmt.Println(len(ch)) // 2
	fmt.Println(cap(ch)) // 3
	ch <- 3
	// ch <- 4 // fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3
	// fmt.Println(<-ch) // fatal error: all goroutines are asleep - deadlock!
	// ###############################

	// Однонаправленные каналы
	// var ch2 chan<- int   // канал только для отправки данных
	// var outCh <-chan int // канал только для получения данных
	// Перепишем factorial

	// Закрытие канала
	ch2 := make(chan int, 5)
	ch2 <- 1
	ch2 <- 2
	close(ch2)
	// ch2 <- 3 // panic: send on closed channel

	fmt.Println(<-ch2) // 1
	fmt.Println(<-ch2) // 2
	fmt.Println(<-ch2) // 0 - default value

	// checking channel closure
	ch3 := make(chan int, 3)
	ch3 <- 10
	ch3 <- 20
	close(ch3)

	// 10 20 Channel closed!
	for i := 0; i < cap(ch3); i++ {
		if val, opened := <-ch3; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed!")
		}
	}

	// Передача потоков данных
	channel := make(chan int)
	go sumTo(5, channel)

	// получаем данные из потока
	for num := range channel {
		fmt.Println(num) // 1 3 6 10 15
	}

	m := make(chan string)
	go getMessage(m)
	fmt.Println(<- m)
}

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}

func sumTo(n int, ch chan int) {
	defer close(ch)
	result := 0
	for i := 1; i <= n; i++ {
		result += i
		ch <- result // посылаем по числу
	}
}

/*
Написать функцию getMessage. Она должна запросить сообщение от пользователя,
а затем передаеть его в канал.

Подготовиться по вопросам.
1. Что такое горутина?
2. Что такое канал?
3. Как горутина работает с небуферизированными каналами?
4. Как горутина работает с буферизированными каналами?
*/

func getMessage(ch chan string) {
	var message string
	fmt.Print("Введите сообщение: ")
	fmt.Fscan(os.Stdin, &message)
	ch <- message
}

// m := make(chan string)
// go getMessage(m)
// fmt.Println(<- m)
