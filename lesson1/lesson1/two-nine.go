package main
 
import "fmt"

func main () {

	var k int
	fmt.Scan(&k)
	s:=(k+1)%12
	fmt.Println(s)
}