//位运算
//原码，反码，补码   计算机以补码进行运算
//-2为例 原码 1000 0010  反码 1111 1101  补码 1111 1110
//&按位与  |按位或  ^按位异或（相同为0，不同为1）

package main
import "fmt"
func main(){
	// var a int = 1 >> 2 //0
	// var b int = -1 >> 2 //-1
	// var c int = 1 << 2 //4
	// var d int = -1 << 2 //-4
	// fmt.Println(a,b,c,d)
	fmt.Println(4|5) //5
	
}