package main
import "fmt"

func main(){
	var x int
	num_array := [10]int{2,5,4,3,6,7,8,9,11,12}
	fmt.Println("Enter the number you want to search")
	fmt.Scan(&x)
	for a := 0; a<10; a++ {
		if x == num_array[a]{
			fmt.Println("Found")
			break
		}
	}

}