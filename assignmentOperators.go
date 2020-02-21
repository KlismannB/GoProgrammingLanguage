package main
import "fmt"

func main(){
	var i int = 10
	
	i += 5
	fmt.Println("i += 5 is ", i)
	// OUTPUT: i += 5 is 15

	i -= 3
	fmt.Println("i -= 3 is ", i)
	// OUTPUT: i -= 3 is 12


	i *= 4
	fmt.Println("i *= 4 is ", i)
	// OUTPUT: i *= 4 is 48


	i /= 2
	fmt.Println("i /= 2 is ", i)
	// OUTPUT: i /= 2 is 24
}