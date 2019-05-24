package main
import "fmt"
import "bufio"
import "os"

func main() {

	var t int
	var sum int = 0
	fmt.Scanf("%d", &t)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for o := 0; o < t; o++ {
	    scanner.Scan()
	    val := toInt(scanner.Bytes())
	    sum += val
	}
	fmt.Println(sum)
}

func toInt(buf []byte) (n int) {
    for _, v := range buf {
        n = n*10 + int(v-'0')
    }
    return
}