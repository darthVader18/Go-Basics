package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	flag := true

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag = false
			break
		} else if i == 1 {
			continue
		}
	}
	fmt.Println(flag)

	day := "Friday"
	switch day {
	case "Friday":
		fmt.Println("TGIF")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday", "Wednesdayy":
		fmt.Println("Boring")
	default:
		fmt.Println("default")
	}

	switch {
	case day == "Friday":
		fmt.Println("TGIF")
		break
	}
}
