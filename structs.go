package main
import "fmt"

type planet struct {
	name string
	galaxyName string
	numberOfMoons int
}

func main(){
	earth := Planet{"Earth","MilkyWay",1}
	fmt.Println(earth)
}
// OUTPUT: {Earth MilkyWay 1}