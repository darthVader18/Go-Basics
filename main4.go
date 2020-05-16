package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("vim-go")
	Anything(1.4455)
	Anything("Chaitanya")
	Anything(struct{}{})

	mymap := make(map[string]interface{})

	mymap["Name"] = "DaRtH VaDeR"
	mymap["Age"] = 18
	fmt.Println(mymap)
}
