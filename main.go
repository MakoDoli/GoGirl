package main

// import ("fmt"; "math"; "math/rand"; "runtime"; "time")
import ("fmt")


func main(){
var taskOne = "1 . Watch Golang course"

var taskItems = []string {"Watch next ep of Homeland", "Make  coffee", "Water plants"}
// fixed size - array, without - slice
	fmt.Println("#### Welcome to our Todo ####")
	fmt.Println() // skip one line
	fmt.Println(taskOne)
	fmt.Println("2 . Practice Golang")
	fmt.Println(3,". Take reward")
	//fmt.Println(taskItems)
	for index, task := range taskItems{
		fmt.Println(index+4, ".", task)
	}
}