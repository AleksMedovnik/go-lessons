package goroutines

import (
	"fmt"
	"sync"
)

var counter int = 0 //  общий ресурс

func Mutexs() {
	ch := make(chan bool)

	// Ограничим одновременный доступ к общему ресурсу
	var mutex sync.Mutex
	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}
	for i := 1; i < 5; i++ {
		<-ch
	}
}

func work(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock() // блокируем доступ к переменной counter
	counter = 0
	for i := 1; i < 5; i++ {
		counter++
		fmt.Println("Goroutine:", i, "-", counter)
	}
	mutex.Unlock() // деблокируем доступ
	ch <- true
}


/* 
var counter int = 0 

func main() {
	ch := make(chan bool)
	for i := 1; i < 5; i++ {
		go work(i, ch)
	}
}

func work(number int, ch chan bool) {
	for i := 1; i < 5; i++ {
		counter++
		fmt.Println("Goroutine:", i, "-", counter)
	}
	ch <- true
}

Сделать так, чтобы получить следующий результат:

Goroutine: 1 - 1
Goroutine: 2 - 2 
Goroutine: 3 - 3 
Goroutine: 4 - 4 
Goroutine: 1 - 5 
Goroutine: 2 - 6 
Goroutine: 3 - 7 
Goroutine: 4 - 8 
Goroutine: 1 - 9 
Goroutine: 2 - 10
Goroutine: 3 - 11
Goroutine: 4 - 12
Goroutine: 1 - 13
Goroutine: 2 - 14
Goroutine: 3 - 15
Goroutine: 4 - 16
*/