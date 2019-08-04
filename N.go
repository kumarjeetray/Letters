/*

**      *
* *     *
*  *    *
*   *   *
*    *  *
*     * *
*      **

*/
package main
import "fmt"
func main(){
	n:=2;
	for i:=1;i<=7;i++{
		for j:=1;j<=9;j++{
			if (j==1 || j==9 || j==n){
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		n++
		fmt.Println()
	}
}
