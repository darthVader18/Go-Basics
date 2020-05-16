package main

import "fmt"

type Car interface {
	Drive()
	Stop()
	Airbags()
}

type Suzuki struct {
	SuzukiModel string
}

func NewModel(arg string) Car {
	return &Suzuki(arg)
}

type Honda struct {
	HondaModel string
}

func (s *Suzuki) Stop() {
	fmt.Println("Stopping Suzuki")
}

func (s *Suzuki) Drive() {
	fmt.Println("Suzuki on the move")
	fmt.Println(SuzukiModel)
}

func (s *Suzuki) Airbags() {
	fmt.Println("Yes")
}

func (h *Honda) Drive() {
	fmt.Println("Honda on the move")
	fmt.Println(HondaModel)
}

func main() {
	s := NewModel("Breeza")
	h := Honda{"Amaze"}
	s.Drive()
	h.Drive()
}
