/*

     **
    *  *
   ******
  *      *
 *        *
 
 */

 package main
import "fmt"
func main(){
	a:=6
	b:=7
	for i:=1;i<=7;i++{
		for j:=1;j<=13;j++{
			if (j==a || j==b || i==3 && j<=b && j>a){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		a--
		b++
		if b==12{
			break
		}
		fmt.Println()
	}
}
