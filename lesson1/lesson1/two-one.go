package main
 
import "fmt"

func main() {
	var s int
	fmt.Scan(&s)
	m:=s/100
	fmt.Printf ("The %d santimeters are %d metrs",s,m)
}