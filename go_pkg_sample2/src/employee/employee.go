package employee

import ()

type Employee struct {
	Name    string
	Salary  float32
	Manager string
}

type EmpRecs struct {
	Emps []Employee
}
