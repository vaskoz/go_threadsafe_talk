package main

import (
	riaken "github.com/riaken/riaken-core"
	"testing"
	"time"
)

// START_TESTS OMIT
func TestRiakenClosure(t *testing.T) {
	t.Parallel()
	addrs := []string{"foo", "bar", "baz"}
	client, someInt := interface{}(nil), 1
	// START_CLOSURE_INITIALIZATION OMIT
	go func() {
		client = riaken.NewClient(addrs, someInt)
	}()
	// END_CLOSURE_INITIALIZATION OMIT
	// START_PARAMETER_MODIFICATION OMIT
	addrs[0], someInt = "vaskoz", 5
	// END_PARAMETER_MODIFICATION OMIT
	time.Sleep(1 * time.Second)
	t.FailNow() // DATA RACE in addrs[0], but NOT someInt
}

func TestRiakenParameter(t *testing.T) {
	t.Parallel()
	addrs := []string{"foo", "bar", "baz"}
	client, someInt := interface{}(nil), 1
	// START_PARAMETER_PASSING OMIT
	go func(addrs []string) {
		client = riaken.NewClient(addrs, someInt)
	}(addrs)
	// END_PARAMETER_PASSING OMIT
	addrs[0], someInt = "vaskoz", 5
	time.Sleep(1 * time.Second)
	t.FailNow() // DATA RACE in addrs[0], but NOT someInt
}

// END_TESTS OMIT
