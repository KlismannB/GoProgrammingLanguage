package main
import "fmt"

func main(){
	var i[5] int =  [5]int{1, 2, 3, 4, 5}
	var slicedI []int = i[1:3]

	fmt.Println("Sliced = ", slicedI)

	slicedI = []int{7,8}

	fmt.Println("Sliced = ", slicedI)
}

/* OUTPUT:

	Sliced = [2 3]
	Sliced = [7 8]

*/