package main
import "fmt"

func main(){
	var hasGirlfriend bool = false
	var isNerd bool = false
	var isSingleAndReadyToMingle bool = true

	if(!hasGirlfriend && !isNerd){
		isSingleAndReadyToMingle = true
		fmt.Println("has Girlfriend = ", hasGirlfriend)
		fmt.Prinln("is Nerd = ", isNerd)
		fmt.Println("is Single and Ready to Mingle = ", isSingleAndReadyToMingle)
	}
}

/* 	OUTPUT:

	has Girlfriend = false
	is Nerd = false
	is Single And Ready to Mingle = true

*/