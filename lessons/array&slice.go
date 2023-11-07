package lessons

import (
	"fmt"
	"reflect"
	"slices"
)

func Arrays() {
	numbers := [5]int{10, -3, 7, -8, 0}
	fmt.Println(numbers[0])

	numbers2 := [5]int{10, -3, 7, -8, 0}
	fmt.Println(numbers == numbers2) // true

	fmt.Println(len(numbers)) // 5

	// copy array
	numbers3 := numbers
	fmt.Println(numbers == numbers3) // true
	numbers3[0] = 1
	fmt.Println(numbers3[0])      // 1
	fmt.Println(numbers[0])       // 10
	fmt.Printf("%p\n", &numbers3) // 0xc0000104e0
	fmt.Printf("%p\n", &numbers)  // 0xc0000104b0

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	for i, elem := range numbers {
		fmt.Println(i, elem)
	}

	for _, elem := range numbers {
		fmt.Println(elem)
	}

}

func Slices() {
	users := []string{"Alex", "Tom", "Boby"}
	fmt.Println(users)

	users = append(users, "Nataly")
	fmt.Println(users[3]) // Nataly

	users2 := append(users, "Jack")
	fmt.Println(users2) // [Alex Tom Boby Nataly Jack]
	fmt.Println(users)  // [Alex Tom Boby Nataly]

	fmt.Println(users2[1:4]) // [Tom Boby Nataly]
	fmt.Println(users2[:4])  // [Alex Tom Boby Nataly]
	fmt.Println(users2[1:])  // [Tom Boby Nataly Jack]

	users = append(users, "Alice", "Masha", "Denis")
	fmt.Println(users) // [Alex Tom Boby Nataly Alice Masha Denis]
	users3 := []string{"David", "Helen"}
	users = append(users, users3...)

	// delete element
	fmt.Println(users) // [Alex Tom Boby Nataly Alice Masha Denis]
	n := 1
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users) // [Alex Boby Nataly Alice Masha Denis]

	// copy by link
	values := []int{1, 2, 3}
	valuesCopy := values
	values[0] = 1000
	fmt.Println(values[0])     // 1000
	fmt.Println(valuesCopy[0]) // 1000

	// copy by value
	foo := []int{1, 2, 3}
	bar := make([]int, len(foo), cap(foo))
	copy(bar, foo)
	foo[0] = 10
	fmt.Println(foo[0]) // 10
	fmt.Println(bar[0]) // 1

	// make
	v := make([]byte, 5, 10)
	fmt.Println(v) // [0 0 0 0 0]

	// module slices
	numbers := []int{0, 33, -10, 8, 33}
	fmt.Println(slices.Max(numbers)) // 33

	fmt.Println(slices.Index(numbers, 33)) // 1
	fmt.Println(slices.Index(numbers, 77)) // -1

	slices.Sort(numbers)
	fmt.Println(numbers) // [-10 0 8 33 33]

	slices.Reverse(numbers)
	fmt.Println(numbers) // [33 33 8 0 -10]



	// capacity
	ns := make([]int, 10)
	fmt.Println(len(ns)) // 10
	fmt.Println(cap(ns)) // 10
	ns = ns[5:8]
	fmt.Println(len(ns)) // 3
	fmt.Println(cap(ns)) // 5

	sl := make([]int, 5)
	fmt.Printf("Capacity before append %d\n", cap(sl)) // 5
	sl = append(sl, 1)
	fmt.Printf("Capacity after append %d\n", cap(sl)) // 10

	sliceBytes := make([]byte, 5)
	fmt.Printf("Capacity before append %d\n", cap(sliceBytes)) // 5
	sliceBytes = append(sliceBytes, 1)
	fmt.Printf("Lenght after append %d\n", len(sliceBytes))   // 6
	fmt.Printf("Capacity after append %d\n", cap(sliceBytes)) // 16

	// передача slice в функцию
	// https://golang-blog.blogspot.com/2020/10/slices-as-arguments-in-golang.html
	abc := []int{1, 2, 3}
	fmt.Println(abc)                            // [1 2 3]
	changeSlice(abc)                            // [1 1 1]
	fmt.Println(abc)                            // [1 1 1]
	change(abc)                                 // [3 3 3 3]
	fmt.Println(abc)                            // [1 1 1]
	fmt.Println(cap(abc))                       // 3
	fmt.Printf("%p\n", &abc)                    // 0xc00008e000
	fmt.Println(reflect.ValueOf(abc).Pointer()) // 824634417272
	abc = append(abc, 2)
	fmt.Println(abc)                            // [1 1 1 2]
	fmt.Println(cap(abc))                       // 6
	fmt.Printf("%p\n", &abc)                    // 0xc00008e000
	fmt.Println(reflect.ValueOf(abc).Pointer()) // 824634458544
	change(abc)                                 // [3 3 3 3 3]
	fmt.Println(abc)                            // [3 3 3 3]
	abc = changeSliceNew(abc)
	fmt.Println(abc) // [4 4 4 4 4]

}

func changeSlice(abc []int) {
	for i := range abc {
		abc[i] = 1
	}
	fmt.Println(abc)
}

func change(abc []int) {
	// fmt.Printf("%p\n", &abc)
	abc = append(abc, 4)
	for i := range abc {
		abc[i] = 3
	}
	fmt.Println(abc)
}

func changeSliceNew(abc []int) []int {
	abc = append(abc, 4)
	for i := range abc {
		abc[i] = 4
	}
	return abc
}


/* Написать функцию findMaxValue, которая принимает слайс с целыми числами
и возвращает максимальное число.
Слайс в задании предусмотрен гарантированно непустой.
Задачу решить без применения стандартной или сторонних библиотек.
Example:
sl := []int{2, 4, 7, -8, 3}
sl2 := []int{10, 4, 7, -8, 3}
sl3 := []int{2, 4, 7, -8, 8}
fmt.Println(findMaxValue(sl)) // 7
fmt.Println(findMaxValue(sl2)) // 10
fmt.Println(findMaxValue(sl3)) // 8
*/

/* Написать функцию findLeftValue, которая принимает целое число и слайс с
целыми числами и возвращает индекс первого (левого) вхождения принимаемого числа в слайсе.
Example:
numbers := []int{0, 33, -10, 8, 33}
fmt.Println(findLeftValue(33, numbers)) // 1

Если слайс пустой или число не найдено, то вернуть -1
Example:
numbers := []int{0, 33, -10, 8, 33}
empty := []int{}
fmt.Println(findLeftValue(56, numbers)) // -1
fmt.Println(findLeftValue(56, empty)) // -1

Задачу решить без применения стандартной или сторонних библиотек. */

/* Написать функцию findRightValue, которая принимает целое число и слайс с
целыми числами и возвращает индекс последнего (первого справа) вхождения принимаемого числа в слайсе.
Example:
numbers := []int{0, 33, -10, 8, 33}
fmt.Println(findRightValue(33, numbers)) // 4

Если слайс пустой или число не найдено, то вернуть -1
Example:
numbers := []int{0, 33, -10, 8, 33}
empty := []int{}
fmt.Println(findRightValue(56, numbers)) // -1
fmt.Println(findRightValue(56, empty)) // -1

Задачу решить без применения стандартной или сторонних библиотек. */




func findMaxValue(arr []int) int {
	x := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > x {
			x = arr[i]
		}
	}
	return x
}

func findLeftValue(n int, arr []int) int {
	x := -1
	for i, elem := range arr {
		if elem == n {
			x = i
			return x
		}
	}
	return x
}

func findRightValue(n int, arr []int) int {
	x := -1
	for i, elem := range arr {
		if elem == n {
			x = i
		}
	}
	return x
}

