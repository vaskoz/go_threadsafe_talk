package main

import (
	"fmt"
	"sync"
	"testing"
)

type SomeStruct struct {
	first   string
	last    string
	address string
	age     int
}

// Copy everything
type SafePeople interface {
	FirstName() string
	LastName() string
	SetFirstName(firstName string)
	SetLastName(lastName string)
	GetStruct() SomeStruct
}

// START_PROTECTED_STRUCT OMIT
type protectedStruct struct {
	mutex      sync.Mutex
	somePerson *SomeStruct
}

// END_PROTECTED_STRUCT OMIT

func NewSafePerson(firstName string, lastName string) SafePeople {
	return &protectedStruct{
		somePerson: &SomeStruct{
			first: firstName,
			last:  lastName,
		},
	}
}

// START_SAFE_STRUCT_TEST OMIT
func TestSafeStructs(t *testing.T) {
	t.Parallel()
	c := make(chan struct{})
	s := SomeStruct{first: "a", last: "b", address: "c", age: 10}
	fmt.Printf("struct pointer: %p\n", &s)
	fmt.Printf("first pointer: %p\n", &(s.first))
	go SaferStructs(c, s)
	s.first += "w"
	<-c
}

// END_SAFE_STRUCT_TEST OMIT

// START_POINTER_STRUCT_TEST OMIT
func TestPointerToStructs(t *testing.T) {
	t.Parallel()
	c := make(chan struct{})
	s := SomeStruct{first: "v", last: "z", address: "d", age: 10}
	fmt.Printf("struct pointer: %p\n", &s)
	fmt.Printf("first pointer: %p\n", &(s.first))
	go DangerousStructPointers(c, &s)
	s.first += "w"
	<-c
	t.FailNow() // Because of DATA RACE
}

// END_POINTER_STRUCT_TEST OMIT

// START_INTERFACE_STRUCTS OMIT
func TestInterfaceToStructs(t *testing.T) {
	c := make(chan struct{})
	safe := NewSafePerson("joe", "smith")
	go func() {
		safe.FirstName()
		c <- struct{}{}
	}()
	safe.SetFirstName("changed")
	<-c
}

// END_INTERFACE_STRUCTS OMIT

// START_SAFE_STRUCT OMIT
// NO DATA RACE because struct gets copied, this doesn't help unless underlying values
// within struct are also copied
func SaferStructs(c chan struct{}, s SomeStruct) {
	fmt.Printf("struct pointer: %p\n", &s)
	fmt.Printf("first pointer: %p\n", &(s.first))
	s.first += "x"
	c <- struct{}{}
}

// END_SAFE_STRUCT OMIT

// START_POINTER_STRUCT OMIT
func DangerousStructPointers(c chan struct{}, s *SomeStruct) {
	fmt.Printf("struct pointer: %p\n", s)
	fmt.Printf("first pointer: %p\n", &(s.first))
	s.first += "x"
	c <- struct{}{}
}

// END_POINTER_STRUCT OMIT

// START_FIRST_NAME OMIT
func (ps *protectedStruct) FirstName() string {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	return ps.somePerson.first
}

// END_FIRST_NAME OMIT

func (ps *protectedStruct) LastName() string {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	return ps.somePerson.last
}

// START_SET_FIRST_NAME OMIT
func (ps *protectedStruct) SetFirstName(firstName string) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	ps.somePerson.first = firstName
}

// END_SET_FIRST_NAME OMIT

func (ps *protectedStruct) SetLastName(lastName string) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	ps.somePerson.last = lastName
}

func (ps *protectedStruct) GetStruct() SomeStruct {
	return *ps.somePerson
}
