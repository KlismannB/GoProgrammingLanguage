package main
import "fmt"

func helloFunction(){
	fmt.Println("Hello World!")
}

func add(i int, j int){ // Pass the parameters with the variable and the type of it
	return i + j
}

func mult(i, j){
	return i * j
}

func main(){
	helloFuction()
	i := 2
	j := 3
	resultSum := add(i, j)
	fmt.Println("The sum result = %v", resultSum)
	// OUTPUT: The sum result = 5
	resultMult := mult(i, j)
	fmt.Println("The multiplication result = %v", resultMult)
	// OUTPUT: The multiplication resul = 6
}