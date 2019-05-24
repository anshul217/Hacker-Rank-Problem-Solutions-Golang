package main
import "fmt"
import "bufio"
import "os"

func main() {

	var a [3]int
	var b [3]int
	var alice int = 0
	var bob int = 0


	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for o := 0; o < 3; o++ {
	    scanner.Scan()
	    a[o] = toInt(scanner.Bytes())
	}
	for o := 0; o < 3; o++ {
	    scanner.Scan()
	    b[o] = toInt(scanner.Bytes())
	}
	if a[0] > b [0]{
		alice += 1
	}else if a[0] < b[0]{
		bob += 1
	}

	if a[1] > b [1]{
		alice += 1
	}else if a[1] < b[1]{
		bob += 1
	}
	if a[2] > b [2]{
		alice += 1
	}else if a[2] < b[2]{
		bob += 1
	}

	fmt.Println(alice, bob)

}

func toInt(buf []byte) (n int) {
    for _, v := range buf {
        n = n*10 + int(v-'0')
    }
    return
}