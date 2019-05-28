package main
import "fmt"
func main(){
	var num [5]int
	var n int
	var sum[5]int
	for i:=0;i<5;i++{
		fmt.Scan(&n)
		num[i] = n
	}
	sum[0]  = num[0]+num[1]+num[2]+num[3]
	sum[1]  = num[1]+num[2]+num[3]+num[4]
	sum[2]  = num[2]+num[3]+num[4]+num[0]
	sum[3]  = num[3]+num[4]+num[0]+num[1]
	sum[4]  = num[4]+num[0]+num[1]+num[2]
	var min int = sum[0]
	var max int = sum[0]
	for i:=0;i<5;i++{
		if max < sum[i]{
			max = sum[i]
		}
		if min > sum[i]{
			min=sum[i]
		}
	}
	fmt.Println (min,max)
}
