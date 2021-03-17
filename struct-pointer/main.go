package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " S.T"
}

func (student Student) displayStudent() string {
	return fmt.Sprintf("ID : %d\nName : %s\nGPA : %f", student.ID, student.Name, student.GPA)
}

func main() {
	Agus := Student{ID: 1, Name: "Agus", GPA: 4.0}
	fmt.Println(Agus)
	Agus.graduate()
	fmt.Println(Agus.Name)
}
