package main
import "fmt"

func main(){
	var i[5]int
	var slicedI []int = i[1:3] //it means that it will begin by the first number,
							   // in ths case, 1, and goes until the last, 3, 
							   //including the first but excluding the last

	fmt.Println(slicedI)
}

// OUTPUT: [0,0]
// The output is 0 because we haven't determined witch one would 
//be the elements of the array