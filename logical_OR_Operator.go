package main
import "fmt"

func main(){
	var hasGirlfriend bool = true
	var isNerd bool = if
	var isSingleAndReadyToMingle bool = true

	if(hasGirlfriend || isNerd){
		isSingleAndReadyToMingle = false
		fmt.Println("has Girlfriend = ", hasGirlfriend)
		fmt.Println("is Nerd = ", isNerd)
		fmt.Println("is Single And Ready to Mingle = ", isSingleAndReadyToMingle)
	}
		/* OUTPUT:

		has Girlfriend = true
		is Nerd = false
		is Single And Ready to Mingle = false

	*/	
}