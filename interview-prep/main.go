package main 


import "fmt"

type Animal struct{
	Name string
	age int
}


// asign functions to animal struct
func(a *Animal) eat() string {
	return a.Name + "is eating"
}

func (a *Animal) sleep() string {
	return a.Name + "is sleeping"
}


// fake inheritance > struct can inherite functions for animal
// example animal > dog struct
type dog struct{
	Animal
}

func test(data []int ,swap chan []int) {
	for a, b := 0, len(data)-1; a < b; a, b = a+1, b-1 {
		data[a], data[b] = data[b], data[a]
	} 
	// send to swap channel
	swap <- data
}
func main(){
	newDog := dog{
		Animal{
			Name: "cutee",
			age: 3,
		},
	}
	fmt.Println("-----------------------------------------------------")

	// fmt.Println(newDog.eat())
	fmt.Println(newDog.sleep())
	fmt.Println("-----------------------------------------------------")
	// loop swap
	swap := make(chan []int)
	data := []int{5, 4, 3, 2, 1} 
 
	// make a go routine 
	go test(data,swap)
	// get data from channel
	x := <- swap
	fmt.Println(x)
}