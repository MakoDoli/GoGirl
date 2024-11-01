package basics

import ("fmt"; "math"; "math/rand"; "runtime"; "time"; "GoGirl/bill")

func add(x int, y int,) int{
	
	return x+y
}
func swap(x, y string) (string, string){
	return x, y
}

func split(sum int) (x,y int){
	x= sum *4/9
	y=sum -x
	return
}

func ifElse(a,b, max int) int{
	if v:=a+b; v < max{
		return v
	}else {
		fmt.Printf("%g is more than %g\n", v, max)
	}
	return max
}

func goodDay(){
	t := time.Now()
	switch {
	case t.Hour()<12 :
		fmt.Println("Good morning")
	case t.Hour() <17 :
		fmt.Println("Good afternoon")
	default: 
	fmt.Println("Good evening")
	}
}

func weekDay(){
	fmt.Print("When is Saturday? ")
	today := time.Now().Weekday()
	switch time.Saturday {
case today + 0 :
	fmt.Println( "today")
case today +1 :
	fmt.Println("tomorrow")
case today +2 :
	fmt.Println(" in 2 days")
	default : 
	fmt.Println ("too far away")
}
}
var c, python, java bool

var k, j int =1,2

func printOS(){
	fmt.Print("Go runs on: ")
	switch os:=runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux" :
			fmt.Println("Linux")
			default :
			fmt.Printf("%s \n", os)
		
	}
}

func main2(){

	fmt.Print("My favorite number is", rand.Intn(10));
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Print(add(3,5))

	a, b := swap("I am a", "I am b")
	fmt.Println(a,b)
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	var c, ruby, react = true, false, "no!"
	fmt.Println(k,j, c, ruby, react)

	sum := 0
	for i := 0; i <7; i++{
		sum += i
		fmt.Println(sum)

	}
	fmt.Println(ifElse(2,3, 10), ifElse(6,7,11))
	printOS()
	weekDay()
	goodDay()

	myBill := bill.NewBill("mako's bill")
	fmt.Println(myBill)
}