//获取用户终端输入
package main
import "fmt"
func main(){
	var a byte
	//var b = 2
	//a,b = b,a
	fmt.Println("年龄")
	fmt.Scanln(&a)
	fmt.Println("年龄是",a)
}