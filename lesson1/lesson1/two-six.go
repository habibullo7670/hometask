package main
 
import "fmt"

func main () {

	var h int
	fmt.Scan (&h)
	hh:=h/3600
	mn:=(h%3600)/60
	sc:=h%60
	fmt.Printf ("%d hours %d minuts %d seconds",hh,mn,sc)
}