package main

import "fmt"

func todo() {
	arr := []int{1, 2, 3, 4}
	arr1 := []string{"Hi", "my", "name"}
	arr1 = append(arr1, "is Chaitanya")
	fmt.Println(arr, arr1)
}

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

type Car struct {
	Name         string
	Company      string
	Type         string
	SeatCapacity int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	todo()

	m1, m2 := 2, 5
	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)

	c := Car{
		Name:         "Amaze",
		Company:      "Honda",
		Type:         "Sedan",
		SeatCapacity: 5,
	}
	c.Print()
	c.GetName()
	fmt.Println(c.GetName())
}
