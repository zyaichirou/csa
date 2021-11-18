//@Title		lv1.go
//@Description	使用channel select来实现子协程之间的执行顺序
//@Author		zy
//@Update		2021.11.16

package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup				//保证main函数等待所有协程执行完

func main() {
	var counter = 10				//打印次数
	ach := make(chan struct{},1)	//ach channel 缓冲区长度为1，只需要执行一个a goroutine 用作唤醒a子协程
	bch := make(chan struct{},1)	//bch channel 缓冲区长度为1，只需要执行一个b goroutine 用作唤醒b子协程
	cch := make(chan struct{},1)	//cch channel 缓冲区长度为1，只需要执行一个c goroutine 用作唤醒c子协程
	w.Add(3)					//设置3个计数器

	ach <- struct{}{}				//先对ach发送信号
	go a(ach,counter,bch)			//开启a goroutine
	go b(bch,counter,cch)			//开启b goroutine
	go c(cch,counter,ach)			//开启c goroutine

	w.Wait()						//等待a b c goroutine结束
}



//@title		a(ach struct{}, counter int, bch chan struct{})
//@description	通过chan实现协程间的执行顺序 在ach收到信号时，给bch通道发送信号，同时本身由于ach通道为空导致for语句一直处于空循环
//@author		zy
//@param		ach chan struct{}  counter int  bch chan struct{}
//@return
func a(ach chan struct{}, counter int, bch chan struct{}){
	atimes := 0
	for atimes < counter {		//执行次数
		select {				//判断通道状态
		case <-ach:				//从ach中接收信号
			fmt.Print("A")	//打印A
			atimes++
			bch <- struct{}{}	//给bch通道发送信号
		}
	}
	w.Done()					//本协程执行完成 调用Done()使计数器--
}

//@title		b(bch struct{}, counter int, cch chan struct{})
//@description	通过chan实现协程间的执行顺序 在bch收到信号时，给cch通道发送信号，同时本身由于bch通道为空导致for语句一直处于空循环
//@author		zy
//@param		bch chan struct{}  counter int  cch chan struct{}
//@return
func b(bch chan struct{}, counter int, cch chan struct{}){
	btimes := 0
	for btimes < counter {		//执行次数
		select {				//判断通道状态
		case <-bch:				//从bch中接收信号
			fmt.Print("B")	//打印B
			btimes++
			cch <- struct{}{}	//给cch通道发送信号
		}
	}
	w.Done()					//本协程执行完成 调用Done()使计数器--
}

//@title		c(cch struct{}, counter int, ach chan struct{})
//@description	通过chan实现协程间的执行顺序 在cch收到信号时，给ach通道发送信号，同时本身由于cch通道为空导致for语句一直处于空循环
//@author		zy
//@param		cch chan struct{}  counter int  ach chan struct{}
//@return
func c(cch chan struct{}, counter int, ach chan struct{})  {
	ctimes := 0
	for ctimes < counter {		//执行次数
		select {				//判断通道状态
		case <-cch:				//从cch中接收信号
			fmt.Print("C ")	//打印C
			ctimes++
			ach <- struct{}{}	//给ach通道发送信号
		}
	}
	w.Done()					//本协程执行完成 调用Done()使计数器--
}