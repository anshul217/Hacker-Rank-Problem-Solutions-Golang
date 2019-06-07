package main
import "fmt"
import "strings"
import "strconv" 

func main(){
	var inputDate string
	fmt.Scanf("%s", &inputDate)
	result := strings.Split(inputDate, ":")
	if result[2][2:] == "AM"{
		if result[0] == "12"{
			
			fmt.Println("00:"+result[1]+":"+result[2][:2])
		
		}else{
		fmt.Println(result[0]+":"+result[1]+":"+result[2][:2])
	}
	}else{
		var hh, err = strconv.Atoi(result[0])
		if err == nil{
			if result[0] == "12"{
							fmt.Println(strconv.Itoa(hh)+":"+result[1]+":"+result[2][:2])

			}else{

			fmt.Println(strconv.Itoa(hh+12)+":"+result[1]+":"+result[2][:2])
		}
	}
		
	}
}