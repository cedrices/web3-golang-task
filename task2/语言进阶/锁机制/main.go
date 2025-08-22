package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	counter := 0
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				lock.Lock()
				// 假设这里有一个共享计数器
				counter++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("共享计数器的值:", counter)

	//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var atomicCounter int64 = 0
	var wg1 sync.WaitGroup
	wg1.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg1.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
		}()
	}
	wg1.Wait()
	fmt.Println("无锁计数器的值:", atomic.LoadInt64(&atomicCounter))
}
