//流程控制


package main
import "fmt"
func main(){
	var score int
	fmt.Println("请输入分数：")
	fmt.Scanln(&score)
	if score >= 60{
		fmt.Println("及格")
	}else{//else不能换行
		fmt.Println("不及格")
	}
	// if score < 60 {
	// 	fmt.Println("不及格")
	// }
	
}