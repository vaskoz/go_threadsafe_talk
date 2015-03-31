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

func BenchmarkFmtPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Print("")
	}
}

func BenchmarkMutexFmtPrint(b *testing.B) {
	var mutex sync.Mutex
	for i := 0; i < b.N; i++ {
		mutex.Lock()
		fmt.Print("")
		mutex.Unlock()
	}
}

func BenchmarkUnbufferedChannelFmtPrint(b *testing.B) {
	c := make(chan string)
	for i := 0; i < b.N; i++ {
		go func() { c <- "" }()
		fmt.Print(<-c)
	}
}

func BenchmarkBufferedChannelFmtPrint(b *testing.B) {
	c := make(chan string, 10)
	for i := 0; i < b.N; i++ {
		go func() { c <- "" }()
		fmt.Print(<-c)
	}
}

func BenchmarkConcurrentRWMutexFmtPrint(b *testing.B) {
	var rwMutex sync.RWMutex
	// Load contention generation only
	for i := 0; i < 1000000; i++ {
		go func() {
			rwMutex.RLock()
			rwMutex.RUnlock()
		}()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rwMutex.Lock()
		fmt.Print("")
		rwMutex.Unlock()
	}
}

func BenchmarkConcurrentMutexFmtPrint(b *testing.B) {
	var mutex sync.Mutex
	// Load contention generation only
	for i := 0; i < 1000000; i++ {
		go func() {
			mutex.Lock()
			mutex.Unlock()
		}()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mutex.Lock()
		fmt.Print("")
		mutex.Unlock()
	}
}

func BenchmarkParallelMutexFmtPrint(b *testing.B) {
	var mutex sync.Mutex
	// Load contention generation only
	for i := 0; i < 1000000; i++ {
		go func() {
			mutex.Lock()
			mutex.Unlock()
		}()
	}
	b.ResetTimer()
	b.SetParallelism(10000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mutex.Lock()
			fmt.Print("")
			mutex.Unlock()
		}
	})
}

func BenchmarkParallelChannelFmtPrint(b *testing.B) {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-c:
				fmt.Print(s)
			}
		}
	}()

	b.ResetTimer()
	b.SetParallelism(10000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c <- ""
		}
	})
}
