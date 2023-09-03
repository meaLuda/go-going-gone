package main


// Arrays Slices Maps Range
import (
	print "fmt" 
)

func learnSlices(){
	// Slices

	// empty slice = nil and have a length of 0
	var s []string // a slice/list of string types
	print.Println(" : ",s,s==nil,len(s)==0) //basic checks on the slice
	
	// we can use make to create an non-zero len slice
	var names = make([]string, 3) // size will take only three values.

	print.Println("Slice with length: ",names, " len: ",len(names),"Capacity: ",cap(names))
	// setting and getting values.
	names[1] = "James"
	print.Println(names)


	// Slices can be mutlidimentional
	twoDSlice := make([][]int, 3)

	for i := 0; i < 3; i++{
		innerLen := i + 1
		// loop each index and make a new slice of length=index+1
		twoDSlice[i] = make([]int, innerLen) // create a new slice in main slice at index ii
		// loop and fill again
		for j := 0;j < innerLen; j++{
			twoDSlice[i][j] = i+j // fill in sum of indexes
		}
	}

	print.Println("2d: ",twoDSlice)
}


func learnMaps(){
	// Sometimes called hashes or dictionaries in python.
	// create empty map
	studentAndAge := make(map[string]int)

	// fill values inside
	studentAndAge["Jane"] = 18
	studentAndAge["Peter"] = 12
	studentAndAge["Sam"] = 20

	print.Println("Map of Students:", studentAndAge,"Length: ",len(studentAndAge))

	// Removing a value
	delete(studentAndAge,"Jane")
	print.Println("Map of Students - Jane:", studentAndAge,"Length: ",len(studentAndAge))

}

func leanRanges(){
	// used to iteration on diff data types

	// Loop of lists
	nums := []int{1,2,3,4,5,6,7}
	sum := 0
	print.Println("Old Sum:",sum)

	for _,num := range nums{
		// increase sum 
		sum += num
	}
	print.Println("New Sum:",sum)


	// loop of maps
	myFruits := map[string]string{"a":"Apples","b":"Bananas"}

	for k,v := range myFruits{
		print.Printf(" %s --> %s \n",k,v)
	}
}

func main(){
	// 01.. Slices
	// learnSlices()

	// 02.. Maps
	// learnMaps()

	// 03...Range
	leanRanges()
}