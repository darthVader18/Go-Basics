package main

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Person struct {
	name   string
	height float64
	gender string
	age    int
}

func (p Person) walk() (string, float64) {
	ret_val_speed := "0 kmph"
	ret_val_rating := 0.00
	if p.name == "Chaitanya" {
		ret_val_speed = "90 kmph"
		ret_val_rating = 5.00
	} else if p.name == "Jai" {
		ret_val_speed = "10 kmph"
		ret_val_rating = 1.00
	}
	return ret_val_speed, ret_val_rating
}

func (p Person) show() {
	log.WithFields(log.Fields{
		"Name":   p.name,
		"Height": p.height,
		"Gender": p.gender,
		"Age":    p.age}).Info("Person structure printed..")

	log.Println("Name: ", p.name)
	log.Println("Height: ", p.height)
	log.Println("Gender: ", p.gender)
	log.Println("Age: ", p.age)
}

func main() {

	f, err := os.OpenFile("syslog.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Error("error opening log file: %v", err)
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	//log.SetOutput(f)

	c := Person{
		name:   "Chaitanya",
		height: 5.9,
		gender: "Male",
		age:    19,
	}

	speed, rating := c.walk()
	log.Printf("Speed of %s is %s and rating is %f\n", c.name, speed, rating)
	c.show()

	j := Person{
		name:   "Jai",
		height: 5.9,
		gender: "Female",
		age:    13,
	}

	speed2, rating2 := j.walk()
	log.Printf("Speed of %s is %s and rating is %f\n", j.name, speed2, rating2)
	j.show()

}

func Setuplog() {
	file, err := os.OpenFile("applog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
}
