package gopkg

import "fmt"

const VERSION = 4.0

// 包内的标识符，首字母大小写，用于访问控制！小写只对包内使用，大写包外可使用
const name string = "haha"
const Name string = "hello"

func printName() { // 包内可用，包外不可用
	fmt.Println("func printName", name)
}

func PrintName() {
	fmt.Println("PrintName", name)
}

// init函数 建议每个包 只定义一个
func init() { // init函数用于初始化包使用，无返回值，init函数在import时自动调用
	printName()
	fmt.Println("github gopkg init")
}
