package main
import "fmt"

func main(){
	var i[3] int
		i[0] = 42
		i[1] = 23
		i[2] = 17

	/* 
		Another way to intialize an array is
		i := []int{42, 23, 17}
	*/

	for j:= 0; j < 3; j++{
		fmt.Println("%v˚ element = %v",j, i[j])
	}

}

/* OUTPUT:

	1˚ element = 42
	2˚ element = 23
	3˚ element = 17

*/