package main
import "fmt"
func main(){
	var a int
	//var b = 2
	//a,b = b,a
	fmt.Println("请输入年龄")
	fmt.Scanln(&a)
	fmt.Println("年龄是",a)
}