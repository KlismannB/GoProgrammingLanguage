package main
import "fmt"

func main(){
	var number float32 = 23.17
	fmt.Printf("Data: %v, Type: %T", number, number)

	// OUTPUT: Data: 23.17, Type: float32

	var number1 int =  int(number)
	fmt.Printf("Data: %v, Type: %T", number1, number1)

	// OUTPUT: Data: 23, Type: int
}