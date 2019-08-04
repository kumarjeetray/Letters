/*

**     **
* *   * *
*  * *  *
*   *   *
*       *
*       *
*       *


*/
package main
import "fmt"
func main(){
	n:=2;
	n2:=8
	for i:=1;i<=7;i++{
		for j:=1;j<=9;j++{
			if (j==1 || j==9 || j==n||j==n2){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		if i<4{
		n++
		n2--
		}else{
			n=0
			n2=0
		}
		fmt.Println()
	}
}
