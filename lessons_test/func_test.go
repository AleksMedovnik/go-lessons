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
