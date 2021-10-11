//switch...case...


package main
import "fmt"
func main(){
	var n1 int32 = 1
	var n2 int32 = 2
	var n3 int32 = 3
	switch n1 + n2 {
	case n2 : //n1,n2,n3数据类型需一致
		fmt.Println("ok1")
	case n3, 5, 10 :
		fmt.Println("ok2")
		fallthrough //穿透，继续执行下一层，输出ok3 
	case 4 :
		fmt.Println("ok3")

	default :
		fmt.Println("wrong!")
	}
}