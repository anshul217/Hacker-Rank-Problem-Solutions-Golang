package main
import "fmt"

func main(){
	var length int
	var positive int = 0
	var neg int = 0
	var zeros int =0
	var tmp int

	fmt.Scan(&length)
	for i:=0; i< length; i++{
		fmt.Scan(&tmp)
		if tmp > 0{
			positive += 1
		} else if tmp < 0{
			neg += 1
		}else{
			zeros +=1
		}
	}

	fmt.Printf("%.6f \n",float64(float64(positive)/float64(length)))
	fmt.Printf("%.6f \n",float64(float64(neg)/float64(length)))
	fmt.Printf("%.6f \n",float64(float64(zeros)/float64(length)))

}