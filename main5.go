package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	f := true
	flag := &f

	if flag == nil {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	arr := []string{"Chaitanya", "Singh", "Bisht"}

	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}

	mymap := make(map[string]interface{})
	mymap["Name"] = "Chaitanya"
	mymap["Age"] = 18

	for k, v := range mymap {
		fmt.Printf("Key: %s and value: %v", k, v)
	}
}
