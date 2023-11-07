package lessons

import (
	"fmt"
	"strings"
	"strconv"
)

func Str() {
	fmt.Print("Hello\n")
	fmt.Println(`C:\users\user\go`) // C:\users\user\go

	message := "Hello"
	fmt.Println(message[0])        // 72
	fmt.Printf("%c\n", message[0]) // H

	a := 'a'
	fmt.Println(a)        // 97
	fmt.Printf("%c\n", a) // a

	a += 1
	fmt.Printf("%c\n", a) // b

	// split
	strSlice := strings.Split("Hello, world!", ", ")
	fmt.Println(strSlice) // [Hello world!]

	// join
	fmt.Println(strings.Join(strSlice, ", ")) // Hello, world!

	// Upper Case
	upperLet := fmt.Sprintf("%c", message[0])
	sliceStr := message[1:]
	fmt.Println(upperLet + sliceStr) // Hello

}


/* 
Написать функцию rle(str *string) string
Example:
str := ""
str2 := "abc"
str3 := "abbccc"
str4 := "abbc"
str5 := "aab"
str6 := "aaaaaaaaaab"
fmt.Println(rle(&str))  // ""
fmt.Println(rle(&str2)) // abc
fmt.Println(rle(&str3)) // ab2c3
fmt.Println(rle(&str4)) // ab2c
fmt.Println(rle(&str5)) // a2b
fmt.Println(rle(&str6)) // a10b
*/

func rle(str *string) string {
	if len(*str) == 0 {
		return *str
	}
	strSlice := strings.Split(*str, "")
	resultSlice := []string{}
	resultSlice = append(resultSlice, strSlice[0])
	count := 1

	for i := 1; i < len(strSlice); i++ {
		if strSlice[i] == resultSlice[len(resultSlice)-1] {
			count++
			if i == len(strSlice)-1 {
				resultSlice = append(resultSlice, strconv.Itoa(count))
			}
		} else {
			if count > 1 {
				resultSlice = append(resultSlice, strconv.Itoa(count), strSlice[i])
			} else {
				resultSlice = append(resultSlice, strSlice[i])
			}
			count = 1
		}
	}
	return strings.Join(resultSlice, "")
}