package main
import "fmt"

func timeTravel(){
	fmt.Println("You are beyond the main. PS: Imagine The Hulk saying: Time Travel!")
}

func main(){
	defer timeTravel()
	fmt.Println("The main ends here.")
}

/* OUTPUT:
	The main ends here.
	You are beyond the main. PS: Imagine The Hulk saying: Time Travel!
*/