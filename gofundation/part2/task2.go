package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func Task2Run1() {
	wg := sync.WaitGroup{}
	//lock := sync.Mutex{}
	wg.Add(2)
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Printf("r1:%v \n", i)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Printf("r2:%v \n", i)
			}
		}
		wg.Done()
	}()
	wg.Wait()

}

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func TaskRun2(tasks []func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	durations := []time.Duration{}
	for idx := range tasks {
		go func(task func()) {
			fmt.Printf("func value:%v \n", &task)
			start := time.Now()
			defer wg.Done()
			task()
			durations = append(durations, time.Since(start))
		}(tasks[idx])
	}
	wg.Wait()
	var sum time.Duration = 0
	for ix := range durations {
		sum += durations[ix]
	}
	fmt.Printf("total cost : %v \n", sum)
}

func main() {
	//TaskRun1()

	tasks := []func(){}
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() {
			r := rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
		})
	}
	TaskRun2(tasks)
}
