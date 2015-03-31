package main

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkUnbufferedChannel(b *testing.B) {
	c := make(chan int)
	for i := 0; i < b.N; i++ {
		go func(val int) { c <- val }(i)
		<-c
	}
}

func BenchmarkBufferedChannel(b *testing.B) {
	c := make(chan int, 10)
	for i := 0; i < b.N; i++ {
		go func(val int) { c <- val }(i)
		<-c
	}
}

func BenchmarkMutex(b *testing.B) {
	var mutex sync.Mutex
	for i := 0; i < b.N; i++ {
		mutex.Lock()
		mutex.Unlock()
	}
}

func Benchmark90R10WMutex(b *testing.B) {
	var rwMutex sync.RWMutex
	for i := 0; i < b.N; i++ {
		rwMutex.RLock()
		rwMutex.RUnlock()
		if i%10 == 0 {
			rwMutex.Lock()
			rwMutex.Unlock()
		}
	}
}

func BenchmarkFmtPrintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Print("")
	}
}

func BenchmarkMutexFmtPrintln(b *testing.B) {
	var mutex sync.Mutex
	for i := 0; i < b.N; i++ {
		mutex.Lock()
		fmt.Print("")
		mutex.Unlock()
	}
}
