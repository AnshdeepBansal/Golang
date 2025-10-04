package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var name string
	name = "anshdeep"
	fmt.Println(name)
	var arr [2]int
	arr[0] = 2
	fmt.Println(arr)
	//create a map in golang and add some values to it te=hen print it
	map1 := make(map[string]int)
	map1["anshdeep"] = 23
	map1["john"] = 30
	fmt.Println(map1)
}
