package main
import "fmt"

func main(){
	var number int = 42 // Is the same as declaring and after assign the value to it's variable
	fmt.Printf("Data: %v, Type: %T", number, number)

	// OUTPUT: Data: 42, Type: int

	var fNumber float32 = float32(number)
	fmt. Printf("Data: %v, Type: %T", fNumber, fNumber)

	// OUTPUT: Data: 42, Type: float32
}