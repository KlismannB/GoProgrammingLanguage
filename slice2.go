package main
import "fmt"

func main(){
	i := []int{8, 3, 12, 6, 10}

	var slicedI = []int = i[1:3]

	fmt.Println("Sliced I = ",slicedI)
	fmt.Println("i = ", i)
}

//OUTPUT: Sliced I = [3 12]
//        i = [8 3 12 6 10]