package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前时间
	timeObj := time.Now()
	fmt.Println(timeObj)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	// 02中的2表示宽带，0表示不足2位以0填充

	fmt.Println("------------------------------------")
	// 日期格式化
	/**
	golang中
	2006 年
	01 月
	02 日
	03 时   03代表12小时制  15代表24小时制
	04 分
	05 秒
	*/
	timeStr := timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)

	fmt.Println("------------------------------------")
	// 获取当前毫秒时间戳
	unixTime := timeObj.Unix()
	fmt.Println(unixTime)
	// 获取当前纳秒时间戳
	unixNanoTime := timeObj.UnixNano()
	fmt.Println(unixNanoTime)
	// 时间戳转time
	timeObj1 := time.Unix(int64(unixTime), 0) // 第一个参数为毫秒  第二个参数为纳秒
	fmt.Println(timeObj1)
	// 日期字符串转time
	timeOnj2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-08-03 21:49:59", time.Local)
	fmt.Println(timeOnj2)
	// time包的常量
	millisecond := time.Millisecond
	second1 := time.Second
	fmt.Println(millisecond)
	fmt.Println((second1))
	fmt.Println("------------------------------------")
	// time定时器
	timeTicker := time.NewTicker(time.Second) // 入参间隔时间
	for c := range timeTicker.C {
		fmt.Println(c)

		timeTicker.Stop() // 停止执行，一定要调，不然在内存中不回销毁
	}

	// 另外一种方式通过休眠+死循环
	time.Sleep(time.Second)
	for {
		time.Sleep(time.Second)
	}
}
