package main
import "fmt"

func main(){
	const PI float32 = 3.1415 
	/*  Using const keyword it makes imposible to change 
		the value of the variable, that in some cases is
		pretty much important, because avoid, for example,
		other programmers to change the name of the company,
		as I will implement.	
	*/
	fmt.Printf("Data: %v, Type: %T", PI, PI)
	// OUTPUT: Data: 3.1415, Type: float32
	
	const COMPANY_NAME string = "The Most Incredible Company of the World"
	fmt.Printf("Data: %v, Type: %T", COMPANY_NAME, COMPANY_NAME)
	// OUTPUT: Data: The Most Incredible Company of the World, Type: string

	/* If the code were wrote like this:

	package main
	import "fmt"

	func main(){
		var COMPANY_NAME string = "TMICotW"

		fmt.Printf("Data: %v", COMPANY_NAME)
		// OUTPUT: Data: TMICofW

		COMPANY_NAME = "The Worst Company"

		fmt.Printf("Data: %v", COMPANY_NAME)
		// OUTPUT: Data: The Worst Company
	}

	*/

}