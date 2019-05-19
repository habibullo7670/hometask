package main
 
import "fmt"

func main () {

	var k,n int
	fmt.Scan (&k,&n)
	r:=(k+n-1)%7
	fmt.Println (r)
}