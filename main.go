package main

// import ("fmt"; "math"; "math/rand"; "runtime"; "time")
import ("fmt"; "net/http")

var taskOne = "1 . Watch Golang course"
var taskItems = []string {"Watch next ep of Homeland", "Make  coffee", "Water plants"}
func main(){
http.HandleFunc("/",  helloUser  )
http.HandleFunc("/tasks",  showTasks  )
http.HandleFunc("/name", yourName)
http.ListenAndServe(":8080", nil)


///////////////////////////////





// fixed size - array, without - slice
	fmt.Println("#### Welcome to our Todo ####")
	fmt.Println() // skip one line
	fmt.Println(taskOne)
	fmt.Println("2 . Practice Golang")
	fmt.Println(3,". Take reward")
	taskItems=addTask(taskItems, "Go for a walk")
	printTasks(taskItems)
	
}

func yourName(writer http.ResponseWriter, request *http.Request){

	var userName string


fmt.Fprintln(writer, "Enter your name")
fmt.Scan(&userName)
fmt.Fprintf(writer, "User's name is %v \n", userName)
fmt.Println(userName)
}

func showTasks(writer http.ResponseWriter, request *http.Request){

	for index, task:= range taskItems {
		fmt.Fprintln(writer, index,task)
	}

}

func helloUser(writer http.ResponseWriter, request *http.Request){
	greeting := "Hello user, Welcome to Go!"
	
	fmt.Fprintln(writer, greeting)
	}

func printTasks(taskItems []string) {
	for index, task := range taskItems{
		listNumber := index+4
		fmt.Printf("%d . %s\n", listNumber, task)
	}
}

func addTask(taskItems []string, newTask string) []string{
taskItems = append(taskItems, newTask)
printTasks(taskItems)
return taskItems
}