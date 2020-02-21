package main
import "fmt"

func main(){
	i := []int{1, 2, 3, 4, 5}
	var slicedI []int = i[1:3]

	var pI *int

	pI = &i[2]

	fmt.Println("Sliced = ", slicedI)
	fmt.Println("i = ", i)

	*pI = 6

	fmt.Println("i = ", i)
	fmt.Println("Sliced = ", slicedI)

	slicedI = i[1:3]

	fmt.Println("i = ", i)
	fmt.Println("Sliced = ", slicedI)
}

/* OUTPUT:

	Sliced = [2 3]
	i = [1 2 3 4 5]

	i = [1 2 6 4 5]
	Sliced = [2 3]

	i = [1 2 6 4 5]
	Sliced = [2 6]

*/