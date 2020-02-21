package main
import "fmt"

func main(){
	/* BITWISE COMPLEMENT OPERATOR */

	var i int = 2
	fmt.Println("^i = ", ^i) // This operator add one to i and invert his signal
	// OUTPUT: ^i = -3

	/* LOGICAL NOT OPERATOR */

	var x bool = true
	fmt.Prinln("!x = ", !x)
	// OUTPUT: !x = false
}