package main

import (
	"fmt"
	"testing"
)

// START_TESTS OMIT
// START_TEST_ARRAYS OMIT
func TestArrays(t *testing.T) {
	t.Parallel()
	data := [...]string{"vasko", "zdravevski", "denver"}
	fmt.Printf("data pointer: %p\n", &data)
	fmt.Printf("element pointer: %p\n", &(data[0]))
	c := make(chan struct{})
	go SafeArrays(c, data)
	data[0] = "notvaskoanymore"
	<-c
}

// END_TEST_ARRAYS OMIT

// START_SAFE_ARRAYS OMIT
func SafeArrays(c chan struct{}, data [3]string) {
	fmt.Printf("data pointer: %p\n", &data)
	fmt.Printf("element pointer: %p\n", &(data[0]))
	data[0] += "x"
	c <- struct{}{}
}

// END_SAFE_ARRAYS OMIT
// END_TESTS OMIT
