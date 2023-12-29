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
