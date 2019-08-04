/*

*******
*      
*      
*******
      *
      *
*******

*/
package main
import "fmt"
func main(){
	n:=1
	b:=5
	for i:=1;i<=7;i++{
		for j:=1;j<=7;j++{
			if (i==1||i==7||i==4||j==1 && i==n||j==7 && i==b){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		if i==b{
			b++
		}
		if (i<3){
			n++
		}else{
			n=0
		}
		fmt.Println()
	}
}
