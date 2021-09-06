package main

import (
	"fmt"
	"os"
)

func main() {
	// 文件读取/写入的三种方式
	/*
		1.os.Open("filePath")
		file.Read() / file.write
		需要关闭文件流
	*/
	file, err := os.Open("")
	if err != nil {
		fmt.Println(file)
		return
	}

	/*
		2.os.Open("filePath")
		bufio.NewReader() / bufio.NewWriter()
		需要刷新flush和close文件
	*/

	/*
		3.ioutil.ReadFile / ioutil.WriteFile()
		// 不需要flush和close文件
	*/

	// 判断文件是否读取完成,会抛出一个异常io.EOF

}
