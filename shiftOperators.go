package main
import "fmt"

func main(){

	// SHIFT OPERATORS

	var i uint = 4
	var j uint = 2

	fmt.Println("i << j", i<<j)
	// OUTPUT: i << j = 16
	fmt.Println("i >> j", i>>j)
	// OUTPUT: i >> j = 1
}