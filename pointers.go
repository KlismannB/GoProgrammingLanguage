package main
import "fmt"

func main(){
	message := "Hello Universe"

	var pMessage *string  // I'm creating a pointer using the * and defining the type of the pointer 
	pMessage = &message
	fmt.Println("Message is = ", *pMessage)
	fmt.Println("Message Address is =" , pMessage)
	/* OUTPUT:
		Message is = Hello Universe
		Message Address is = SOMEADRESS
	*/
	*pMessage = "Hello Cosmos"
	fmt.Println("Message is = ", *pMessage)
	fmt.Println("Message Address is = ", pMessage)
	/* OUTPUT:
		Message is = Hello Cosmos
		Message Address is = SOMEADRESS
	*/
}