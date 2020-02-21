package main
import "fmt"

func main(){
	var weight float32 // It can be a float64 too, because depends on the size of the bits or 32 bits or 64 bits
	weight = 17.32
	fmt.Printf("Data: %v, Type: %T", weight, weight)
}

// OUTPUT: Data: 17.32, Type: float32