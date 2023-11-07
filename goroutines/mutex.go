package goroutines

import (
	"fmt"
	"sync"
)

var counter int = 0 //  общий ресурс

func Mutexs() {
	ch := make(chan bool)

	/* for i := 1; i < 5; i++ {
		go work(i, ch)
	}
	// ожидаем завершения всех горутин
	for i := 1; i < 5; i++ {
		<-ch
	} */

	// Ограничим одновременный доступ к общему ресурсу
	var mutex sync.Mutex
	for i := 1; i < 5; i++ {
		go work2(i, ch, &mutex)
	}
	for i := 1; i < 5; i++ {
		<-ch
	}
}

func work(number int, ch chan bool) {
	counter = 0 //  переопределяем глобальную переменную
	for i := 1; i < 5; i++ {
		counter++
		fmt.Println("Goroutine:", i, "-", counter)
	}
	ch <- true
}

func work2(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock() // блокируем доступ к переменной counter
	counter = 0
	for i := 1; i < 5; i++ {
		counter++
		fmt.Println("Goroutine:", i, "-", counter)
	}
	mutex.Unlock() // деблокируем доступ
	ch <- true
}
