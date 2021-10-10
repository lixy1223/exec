package main
import "fmt"
func main(){
	var n1 int64 = 10
	var n2 int32 
	var n3 int8 

	n2 = int32(n1) + 20
	n3 = int8(n1) + 127

	fmt.Println(n2,n3)
	
	
}