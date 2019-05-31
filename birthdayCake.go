package main
import "fmt"
func main(){
	
	var n int 
	fmt.Scan(&n)
	var s = make([]int, n)
	var largest int = s[0]
	var counter int = 0
	for i:=0;i<n;i++{
		fmt.Scan(&s[i])
	}
	for i := 1;i<n;i++{
		if s[i]>largest{
		largest = s[i]
		}
	}

	for i:=0;i<n;i++{
		if s[i] == largest{
		counter+=1
	}
	}
	fmt.Println(counter)
}