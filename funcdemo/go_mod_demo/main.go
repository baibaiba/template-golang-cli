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
	// 3.使用包，通过import导入,导入格式
	add := tools.Add(1, 2)
	fmt.Println(add)
}
