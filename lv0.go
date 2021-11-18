//@Title		lv0.go
//@Description	用sync.WaitGroup来实现等待子协程
//@Author		zy
//@Update		2021.11.16


package main

import (
	"fmt"
	"sync"
)

var (
	myres = make(map[int]int, 20)
	//mu sync.Mutex
	wg sync.WaitGroup	//保证main函数等待所有协程执行完
)

func factorial(n int) {
	defer wg.Done()		//推迟调用Done()方法将计数器-1
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//mu.Lock()
	myres[n] = res
	//mu.Unlock()
}

func main() {

	for i := 1; i <= 20; i++ {
		wg.Add(1)	//每执行一个goroutine  就增加计数器的值
		go factorial(i)
	}

	//mu.Lock()
	wg.Wait()	//通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
	//mu.Unlock()
}
