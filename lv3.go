//@Title		lv3.go
//@Description	使用管道，实现一个简单通知关闭的Context函数
//@Author		zy
//@Update		2021.11.16


package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


var (
	wgg sync.WaitGroup		//等待所有子协程
	key1 = "name"
	key2 = key1
)

func main()  {
	wgg.Add(2)		//设置计数器值为2
	//context.Backgrount()返回一个空的Context，用作整个Context树的根节点
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key1, "卷王")  //使用WithValue返回ctx的副本，同时附加一对key-value的键值对
	go juan(valueCtx)				//开启子goroutine
	time.Sleep(5 * time.Second)	//设置10秒的协程运行时间
	fmt.Println("卷不动了！大伙喝口水再继续！")
	cancel()						//结束juan goroutine  juan goroutine结束意味着 sonOfJuan也结束  取消一个Context 该节点下所有的context都取消
	wgg.Wait()						//等待所有协程结束
}

//@title		juan(ctx context.Context)
//@description	创建一个子context并开启一个子goroutine 等待Context根节点的cancel() 同时每隔1s打印 ctx.value "coding中"
//@author		zy
//@param		ctx context.Context
//@return
func juan(ctx context.Context)  {
	defer wgg.Done()				//推迟调用Done(),计数器--
	ctx1, _ := context.WithCancel(ctx)	//生成ctx的子context
	valueCtx1 := context.WithValue(ctx1,key2,"卷中卷")	//使用WithValue返回ctx1的副本，同时附加一对key-value的键值对
	go sonOfJuan(valueCtx1)	//子goroutine中又开启另外一个goroutine
	for {
		select {					//选择通信操作
		case <-ctx.Done():			//等待cancel()
			fmt.Println(ctx.Value(key1), "喝水了！")
			return
		default:					//juan goroutine中，sleep 1s
			fmt.Println(ctx.Value(key1),"coding中！")
			time.Sleep(time.Second)
		}
	}
}

//@title		sonOfJuan(ctx context.Context)
//@description	等待Context根节点的cancel() 同时每隔0.25s打印 ctx.value "coding中"
//@author		zy
//@param		ctx context.Context
//@return
func sonOfJuan(ctx context.Context) {
	defer wgg.Done()				//推迟调用Done(),计数器--
	for {
		select {					//选择通信操作
		case <-ctx.Done():			//等待cancel()
			fmt.Println(ctx.Value(key2), "也喝水了！")
			return
		default:					//sonOfJuan goroutine中，只sleep 0.25s
			fmt.Println(ctx.Value(key2),"coding中！")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
