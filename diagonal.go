package main
import "fmt"

func main() {

	var row int
	var diag1 int = 0
	var diag2 int = 0
	fmt.Scan(&row)
	matrix := make([][]int, row)
	for i := range matrix{
		matrix[i] = make([]int, row)
	}
	for i:=0; i<row; i++{
		for j:=0; j<row; j++{
			fmt.Scan(&matrix[i][j])
			if i ==j {
				diag1 += matrix[i][j]
			}
			if (i+j) == (row -1 ){
				diag2 += matrix[i][j]
			}
		}
	} 
	if (diag1 - diag2) >= 0{
		fmt.Println(diag1 - diag2)
		}else{
			fmt.Println((diag1 - diag2)*-1)
		}
	 
}