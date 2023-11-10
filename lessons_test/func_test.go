package lessons_test

import (
	"testing"

	"github.com/AleksMedovnik/go-lessons/lessons"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	got, err := lessons.Sum(1, 2)
	if err != nil {
		t.Errorf("Error")
	}
	if got != 3 {
		t.Errorf("Sum(1, 2) = %d; want 3", got)
	}

}

// go test -v

// go get github.com/stretchr/testify

func TestSumTestify(t *testing.T) {
	var exp1 byte = 3
	var exp2 byte = 155

	t1, err := lessons.Sum(1, 2)
	assert.Nil(t, err)
	assert.Equal(t, exp1, t1)

	t2, err := lessons.Sum(0, 3)
	assert.Nil(t, err)
	assert.Equal(t, exp1, t2)

	t3, err := lessons.Sum(100, 55)
	assert.Nil(t, err)
	assert.Equal(t, exp2, t3)

}

/* 
1. Написать функцию, возвращающую факториал числа.
Принимаемое и возвращаемые числа должны быть представлены типом int64.
Предусмотреть ошибки. Если аргумент меньше 1, то вывести сообщение ошибки "Invalid input number".
Example:
Factorial(1) // -> 1
Factorial(3) // -> 6
Factorial(5) // -> 120
Factorial(11) // -> 39916800
Factorial(0) // -> Invalid input number, -1

2. Написать тесты для данной функции при помощи пакета testify.
*/


func TestFactorial(t *testing.T){
	var exp1 int64 = 6
	var exp2 int64 = 39916800
	var exp3 int64 = 120

	r1, err := lessons.Factorial(3)
	assert.Nil(t, err)
	assert.Equal(t, exp1, r1)

	r2, err := lessons.Factorial(11)
	assert.Nil(t, err)
	assert.Equal(t, exp2, r2)

	r3, err := lessons.Factorial(5)
	assert.Nil(t, err)
	assert.Equal(t, exp3, r3)
}