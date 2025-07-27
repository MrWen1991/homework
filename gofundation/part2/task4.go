package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//r1()
	r2()
}

/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。
一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func r1() {
	//ch := make(chan int, 10)
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Printf("Producer value:%v \n", i)
			//time.Sleep(50 * time.Millisecond)
		}
		close(ch)
		wg.Done()
	}()

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				break
			}
			time.Sleep(2000 * time.Millisecond)
			fmt.Printf("consumer value:%v \n", v)
		}
		wg.Done()
	}()

	wg.Wait()
}

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func r2() {
	ch := make(chan int, 100)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				break
			}
			fmt.Printf("consumer value:%v \n", v)
		}
		wg.Done()
	}()
	wg.Wait()
}
