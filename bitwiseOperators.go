package main
import "fmt"

func main(){
	// BITWISE COMPARATORS
	var i int = 4
	var j int = 2

	fmt.Println("i & j = ", i&j)
	// OUTPUT: i & j = 0
	fmt.Println("i | j = ", i|j)
	// OUTPUT: i | j = 6
	fmt.Println("i ^ j = ", i^j)
	// OUTPUT: i ^ j = 6
}