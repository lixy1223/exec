//流程控制
//单分支、多分支


package main
import "fmt"
func main(){
	// var score int
	// fmt.Println("请输入分数：")
	// fmt.Scanln(&score)
	// if score >= 60{
	// 	fmt.Println("及格")
	// }else{//else不能换行
	// 	fmt.Println("不及格")
	// }
	// if score < 60 {
	// 	fmt.Println("不及格")
	// }
	var score int
	fmt.Println("请输入分数：")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("奖励一辆BWM")
	}else if score <= 99 && score > 80 {//else不能换行
		fmt.Println("奖励一台iphone13")
	}else if score <= 80 && score >= 60{
		fmt.Println("奖励一个iPad")
	}else{
		fmt.Println("奖励一个大耳巴子")
	}
}