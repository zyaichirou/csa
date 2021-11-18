//@Title		lv2.go
//@Description	并发求1~50000以内的素数，使用同步机制等待所有协程执行完，使用互斥锁防止出现竞态问题
//@Author		zy
//@Update		2021.11.16

package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var (
	num = make([]int,0)		//存储素数
	lock sync.Mutex			//使用互斥锁防止竞态问题
	wc sync.WaitGroup		//保证所有协程执行完
)

//@title		prime(n int) bool
//@description	判断一个数是否为素数
//@author		zy
//@param		n int
//@return		bool
func prime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 1 ; i <= int(math.Sqrt(float64(n))) ; i++ {
		if n % i == 0 && i != 1 {
			return false
		}
	}
	return true
}

//@title		betw(l, r int)
//@description	求l到r之间的所有素数 同时把素数append到num中
//@author		zy
//@param		l, r int
//@return
func betw(l, r int) {
	defer wc.Done()					//延迟调用Done(),计数器--
	for i := l; i < r; i++{
		if prime(i) {
			lock.Lock()				//上锁
			num = append(num,i)
			lock.Unlock()			//解锁
		}
	}
}

func main() {
	start := time.Now()				//开始时间

	//并发	开8个协程
	for i := 1; i <= 50000; i += 6250 {
		wc.Add(1)				//计数器++
		go betw(i,i+6250)
	}
	wc.Wait()	//调用Wait()，等待所有协程结束
	fmt.Println("8 goroutine:")


	//正常递归
	//for i := 1; i <= 50000; i++ {
	//	if prime(i) {
	//		num = append(num,i)
	//	}
	//}
	//fmt.Println("Normal circle loop:")

	fmt.Println(len(num))
	fmt.Println(num)

	end := time.Now()				//结束时间
	fmt.Println(end.Sub(start))		//运行时间
}
