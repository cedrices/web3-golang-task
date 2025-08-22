package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 题目:编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	// printOddAndEven()

	//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	tasks := []func(){
		func() {
			fmt.Println("任务1开始")
			for i := 0; i < 1000000; i++ {
			}
			fmt.Println("任务1结束")
		},
		func() {
			fmt.Println("任务2开始")
			for i := 0; i < 2000000; i++ {
			}
			fmt.Println("任务2结束")
		},
	}
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(t func()) {
			start := time.Now()
			t()
			end := time.Since(start)
			fmt.Printf("任务执行时间: %d毫秒\n", end.Milliseconds())
			wg.Done()
		}(task)
	}
	wg.Wait()
}

func printOddAndEven() {

	odd := make(chan int)
	even := make(chan int)
	go func() {

		for i := 1; i <= 10; i++ {
			if i%2 > 0 {
				odd <- i
			}
		}
		close(odd)
	}()

	for i := range odd {
		fmt.Println("奇数:", i)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				even <- i
			}
		}
		close(even)
	}()

	for i := range even {
		fmt.Println("偶数:", i)
	}

}
