package main

import "fmt"

type Car interface {
	Drive()
	Stop()
	//Airbags()
}

type Suzuki struct {
	SuzukiModel   string
	NumberofTyres int
}

type Honda struct {
	HondaModel    string
	NumberofTyres int
}

func (s *Suzuki) Stop() {
	fmt.Println("Stopping Suzuki")
}

func (s *Suzuki) Drive() {
	fmt.Println("Suzuki on the move")
	fmt.Println(s.SuzukiModel)
}

func (s *Suzuki) Airbags() {
	fmt.Println("Yes")
}

func (h *Honda) Drive() {
	fmt.Println("Honda on the move")
	fmt.Println(h.HondaModel)
}

func NewModel(name string, tyre int) Car {
	u := Suzuki{}
	u.SuzukiModel = name
	u.NumberofTyres = tyre
	return &u
}

func main() {
	//Function Calling
	s := NewModel("Breeza", 4)

	s.Drive()
	//Object creation
	h := Honda{}
	h.HondaModel = "Amaze"
	//h.NumberofTyres = 4
	h.Drive()
	fmt.Println(h.NumberofTyres)

}
