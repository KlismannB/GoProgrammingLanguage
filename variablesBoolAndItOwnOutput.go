package main
import "fmt"

func main(){
	var doesUniverseExists bool // Defining variable doesUniverseExists as bool
	
	doesUniverseExists = true // Assign variable doesUniverseExists as true

	fmt.Printf("Data: %v, Type: %T", doesUniverseExists, doesUniverseExists)
	// For those who know C programming, you can use only %v besides using %d, %b, %i, %f  etc...
	
}

// OUTPUT: Data: true, Type: bool