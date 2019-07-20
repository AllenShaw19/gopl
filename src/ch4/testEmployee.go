package main

import "fmt"

type Employee struct {
	ID       int
	Position string
	Salary   int
}

func EmployeeByID(id int, e []Employee) Employee {
	for _, v := range e {
		if v.ID == id {
			return v
		}
	}
	return Employee{}
}

func main() {
	var dilbert Employee
	dilbert.ID = 2
	dilbert.Position = "Programmer"
	dilbert.Salary = 5000

	var phb Employee
	phb.ID = 1
	phb.Position = "boss"
	phb.Salary = 10000

	e := []Employee{phb, dilbert}

	d := EmployeeByID(2, e)
	d.Salary = 0

	u := EmployeeByID(3, e)
	u.Salary = 1230

	fmt.Printf("%+v\n", d)
	fmt.Printf("%+v\n", u)
}
