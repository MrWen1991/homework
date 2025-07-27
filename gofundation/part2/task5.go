package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
var count int = 0

func syncRun1() {
	//lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				time.Sleep(5 * time.Millisecond)
				//lock.Lock()
				count += 1
				fmt.Printf("routine %v, count:%v \n", i, count)
				//lock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全
*/
func syncRun2() {
	var c uint64 = 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		idx := i
		go func() {
			for j := 0; j < 1000; j++ {
				time.Sleep(5 * time.Millisecond)
				//lock.Lock()
				atomic.AddUint64(&c, 1)
				fmt.Printf("routine %v \n", idx)
				//lock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("count:%v \n", atomic.LoadUint64(&c))
}

func main() {
	//syncRun1()
	syncRun2()
}
