/*
*******
*     *
*     *
*******
*     *
*     *
*******

*/
package main
import "fmt"
func main(){
	for i:=1;i<=7;i++{
		for j:=1;j<=7;j++{
			if (i==1 ||i==4 ||j==1||j==7 || i==7){
				fmt.Print("*")
				}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
