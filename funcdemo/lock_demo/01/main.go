package main

import (
	"fmt"
	"sync"
)

// 1.互斥锁
// 2.读写互斥锁

var wg sync.WaitGroup
var mutex sync.Mutex
var count = 0

func main() {
	wg.Add(1)
	go myRead()

	wg.Add(1)
	go myWrite()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go test()
	}

	wg.Wait()
}

// 如果并发执行以下的方法,多协程之间会竞争count
// 使用 go build -race main.go 编译程序生成可执行文件，然后执行，可以看到竞争情况
/**
==================
WARNING: DATA RACE
Read at 0x000000606908 by goroutine 10:
  main.test()
      /workspace/template-golang-cli/funcdemo/lock_demo/main.go:32 +0x3e

Previous write at 0x000000606908 by goroutine 9:
  main.test()
      /workspace/template-golang-cli/funcdemo/lock_demo/main.go:32 +0x5a

Goroutine 10 (running) created at:
  main.main()
      /workspace/template-golang-cli/funcdemo/lock_demo/main.go:24 +0xcf

Goroutine 9 (finished) created at:
  main.main()
      /workspace/template-golang-cli/funcdemo/lock_demo/main.go:24 +0xcf
==================
4
5
6
7
8
9
10
Found 1 data race(s)
*/

// 解决这种资源竞争可以加互斥锁,加了互斥锁，同一时间只能有一个协程进行操作，其它协程既不能读也不能写
func test() {
	mutex.Lock()
	count++
	fmt.Println(count)
	mutex.Unlock()
	wg.Done()
}

func myRead() {
	fmt.Println("读操作")
	wg.Done()
}

func myWrite() {
	mutex.Lock() // 加互斥锁
	fmt.Println("写操作")
	mutex.Unlock()
	wg.Done()
}
