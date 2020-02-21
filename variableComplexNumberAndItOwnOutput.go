package main
import "fmt"

func main(){
	var complexNumber complex64
	complexNumber =  complex(2,3)
	fmt.Printf("Data: %v, Type: %T", complexNumber, complexNumber)
}

// OUTPUT: Data: (2 + 3i), Type: complex64