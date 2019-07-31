package main
import "fmt"
func main(){
	a:=5
	for i:=1;i<=7;i++{
		for j:=1;j<=7;j++{
			if ((j==1)|| (j==a)){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		if i<4{
			a--
		}else{
			a++
		}
		fmt.Println()
	}
}