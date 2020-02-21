package main
import "fmt"

func main(){

	/* -------- POSTFIX OPERATOR -------- */

	var i int = 10
	i++ // Add 1 to i
	fmt.Println("i++ = ", i)
	
	// OUTPUT: i++ = 11

	i = 10
	i-- // Subtract 1 from i
	fmt.Println("i-- = ", i)

	// OUTPUT: i-- = 9

	/* -------- UNARY OPERATOR -------- */

	i = 10
	fmt.Println("+i = ", +i)
	
	// OUTPUT: +i = 10

	fmt.Println("-i = ", -i)

	// OUTPUT: -i = -10

}