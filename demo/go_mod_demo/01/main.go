package main

import (
	"fmt"

	"github.com/gitpod/mycli/funcdemo/go_mod_demo/tools"
)

func main() {
	// 一个包下的多个文件内的方法定义不能重复

	// golang中自定义个包的逻辑
	// 1.项目中必须有go.mod文件，可以通过 【go mod init 项目名】在项目根目录创建
	// 2.定义一个包文件夹，在文件夹中创建包文件(见本文件下的tools文件夹)；一个包文件下可以创建很多包文件，但是包名需要一致
	// 这样一个自定义包就定义好了
	// 3.使用包，通过import导入,导入格式【module名（见go.mod文件） 自定义包路径】
	add := tools.Add(1, 2)
	fmt.Println(add)
	sum := tools.Sum(1, 2, 4, 5, 6)
	fmt.Println(sum)

	// 第三方包引用逻辑
	// 1. import第三方包，然后写上包引用的代码；这时候是不能运行的，因为第三方包还没有下载
	// 运行 go mod download(将包下载到go path) 或者 go mod vendor(将包下载到项目)，
	// 目前idea在引入第三发包时运行项目会自动帮助我们找到包
	// 文件变化：引入第三方包后go.mod文件中会有require引用，同时会自动生成go.sum文件（自动维护）
}
