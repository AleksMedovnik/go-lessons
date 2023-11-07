package lessons

import (
	"fmt"
	"strings"
)

func Generics() {
	sliceInt := []int{1, 2, 3}
	printSlice(sliceInt)

	sliceStr := []string{"Hello", "Bye", "Hi"}
	printSlice(sliceStr)

	s := []int64{2, -3, 7, 8, 9, -4, 0}

	// Filter
	fs := Filter(s, func(i int64) bool { return i > 0 })
	fs2 := Filter(s, func(i int64) bool { return i%2 != 0 })
	fs3 := Filter(sliceStr, func(i string) bool { return len(i) > 2 })

	fmt.Println(fs)  // [2 7 8 9]
	fmt.Println(fs2) // [-3 7 9]
	fmt.Println(fs3) // [Hello Bye]

	// Map
	ms := Map(s, func(i int64) int64 { return i * i })
	ms2 := Map(sliceStr, func(i string) string { return strings.ToLower(i) })

	fmt.Println(ms)  // [4 9 49 64 81 16 0]
	fmt.Println(ms2) // [hello bye hi]
}

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Map[T any](s []T, f func(T) T) []T {
	var r []T
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}
