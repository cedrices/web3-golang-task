package main

import (
	"fmt"
	"sync"
)

func main() {
	//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	printChan()

	//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	bufferedChan := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	go provider(bufferedChan, &wg)
	go consumer(bufferedChan, &wg)
	wg.Wait()
	fmt.Println("所有数字已处理完毕。")
}

func printChan() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go wirteChan(ch, &wg)
	go readChan(ch, &wg)
	wg.Wait()
}

func wirteChan(ch chan<- int, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	wg.Done()
}

func readChan(ch <-chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		num := <-ch
		fmt.Println("接收到的数字:", num)
	}
	wg.Done()
}

func provider(bufferedChan chan<- int, wg *sync.WaitGroup) {
	for i := 1; i <= 100; i++ {
		chanLen := len(bufferedChan)
		if chanLen == 10 {
			fmt.Println("缓冲区已满，等待消费者处理...")
		}
		bufferedChan <- i
	}
	close(bufferedChan)
	wg.Done()
}

func consumer(bufferedChan <-chan int, wg *sync.WaitGroup) {
	for num := range bufferedChan {
		fmt.Println("消费者接受到的数据:", num)
	}
	wg.Done()
}
