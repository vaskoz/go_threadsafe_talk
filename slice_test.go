package main

import (
	"fmt"
	"testing"
)

// START_TEST_SLICES OMIT
func TestSlices(t *testing.T) {
	t.Parallel()
	data := []string{"foo", "bar", "baz"}
	fmt.Printf("data pointer: %p\n", &data)
	fmt.Printf("element pointer: %p\n", &(data[0]))
	c := make(chan struct{})
	go DangerousSlices(c, data)
	data[0] = "notfoo"
	<-c
	t.FailNow() // Because of DATA RACE
}

// END_TEST_SLICES OMIT

// START_TEST_VARIADIC OMIT
func TestVariadicSlices(t *testing.T) {
	t.Parallel()
	data := []string{"foo", "bar", "baz"}
	fmt.Printf("data pointer: %p\n", &data)
	fmt.Printf("element pointer: %p\n", &(data[0]))
	c := make(chan struct{})
	go DangerousVariadic(c, data...)
	data[0] = "notfoo"
	<-c
	t.FailNow() // Because of DATA RACE
}

// END_TEST_VARIADIC OMIT

// START_TEST_VARIADIC_SEPARATE OMIT
func TestVariadicSeparates(t *testing.T) {
	t.Parallel()
	foo, bar, baz := "foo", "bar", "baz"
	c := make(chan struct{})
	go DangerousVariadic(c, foo, bar, baz)
	foo = "notfoo"
	<-c
}

// END_TEST_VARIADIC_SEPARATE OMIT

// START_DANGEROUS_SLICES OMIT
// Never pass mutable data structures in thread-safe libraries
// You can never protect against race conditions in your own code
func DangerousSlices(c chan struct{}, data []string) {
	fmt.Printf("data pointer: %p\n", &data)
	fmt.Printf("element pointer: %p\n", &(data[0]))
	data[0] += "x"
	c <- struct{}{}
}

// END_DANGEROUS_SLICES OMIT

// START_DANGEROUS_VARIADIC OMIT
// No way to change the values of these once they're passed in
// Spec says a copy of the slice is created
func DangerousVariadic(c chan struct{}, data ...string) {
	fmt.Printf("data pointer: %p\n", &data)
	fmt.Printf("element pointer: %p\n", &(data[0]))
	data[0] += "x"
	c <- struct{}{}
}

// END_DANGEROUS_VARIADIC OMIT
