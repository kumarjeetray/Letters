package main
import "fmt"
func main(){
	a:=1
	b:=10
	for i:=1;i<=5;i++{
		for j:=1;j<=10;j++{
			if (j==a || j==b){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		a++
		b--
		fmt.Println()
	}
}