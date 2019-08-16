/*
*******
*     *
*     *
*******
*  *   
*   *  
*    * 

*/
package main
import "fmt"
func main(){
	a:=
	for i:=1;i<=7;i++{
		for j:=1;j<=7;j++{
			if (i==1 ||i==4 ||j==1||j==7 && i<=4){
				fmt.Print("*")
				}else if (i>4 && j==a){
					fmt.Print("*")
					}else{
				fmt.Print(" ")
			}
		}
		if i>4{
			a++
		}
		fmt.Println()	
		}
}
