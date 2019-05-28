package main
import "fmt"

func main(){
	
	var numbers [10]int

	for i:=0; i<10;i++{
	fmt.Scan(&numbers[i])
	}
	var temp int = 0
	for i:=0;i<10;i++{
		for j:=0;j<9;j++{
			if numbers[j] > numbers[j+1]{
				temp = numbers[j]
				numbers[j]=numbers[j+1]
				numbers[j+1]=temp
				break
			}
		}
	}
	fmt.Println(numbers)
	fmt.Println("Complexity O(n2)")
}