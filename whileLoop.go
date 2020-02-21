package main
import "fmt"

// Unfortunately doesn't have a function called while in go language
// But you can also use a for or a flag

/*  
	For those who never programmed, a flag is a thing that you can use when
	you want to acomplish some task, like if you want to count until you find 
	a certain number
 */

func main(){
	var find bool = false
	// Example with a simple for (NEVER ENDS)
	i := 0;

	for i != 1 {
		fmt.Println("Not found yet!")
	}

	// Example using flag (NEVER ENDS) because I haven't marked a point to finish, because the line is commented
	for i := 0; !find; i++{
		fmt.Println("Not found yet!")
		/*
		if(i == 100){
			find = true
			fmt.Println("Finded")
		}
		*/
	}
}

