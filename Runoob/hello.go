package main   // 包声明
/* package main 表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包 */
// 同一个文件夹下的文件只能有一个包名，否则编译报错
 
// import "fmt"   // 引入包
import (
	"fmt"
	"./myMath"
)
/* fmt 包含了格式化 IO 的函数 */

/* 
 * func main() 是程序开始执行的函数
 * main 函数是每一个可执行程序所必须包含的
 * 一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）
 */
func main() {    // ****** { 不能单独放在一行
	/* fmt.Println() 可以将字符串输出到控制台，并在最后自动增加换行字符 \n */
	fmt.Println("Hello World!")
	fmt.Println(mathClass.Add(1,1))
	fmt.Println(mathClass.Sub(1,1))
}
/* 执行Go程序
 * 打开命令行，并进入程序文件保存的目录中
 * go run hello.go 命令执行代码
 * go build 命令生成二进制文件
 */