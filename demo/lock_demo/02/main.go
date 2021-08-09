package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

//  读写互斥锁
func main() {
	// 场景：同一时间可以有很多人读数据，但是同一时间只能一个人写数据

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
}

func read() {
	mutex.RLock() // 加一个读锁，代表同一时间可以同时读
	fmt.Println("读数据")
	time.Sleep(time.Second * 2)
	mutex.RUnlock()
	wg.Done()
}

func write() {
	mutex.Lock() // 加一个互斥锁，同一时间其它协程既不能读也不能写
	fmt.Println("写数据")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}
