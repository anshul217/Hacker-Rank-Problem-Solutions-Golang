package main
import "fmt"
func main(){
	var n [10]int
	for i:=0;i<10;i++{
	fmt.Scan(&n[i])
	}
	var temp int =0 
	for i:=0;i<9;i++{
		for j:=i+1; j<10;j++{
				if n[i] > n[j]{
					temp = n[i]
					n[i] = n[j]
					n[j] = temp
				}
		}
	}
	fmt.Println(n)
}