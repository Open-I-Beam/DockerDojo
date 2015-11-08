package main

import (
	"fmt"
	"jsonbuilder"
	"os"
)

type Employee struct {
	Name   string
	Salary float32
}

type EmpRecs struct {
	emps []Employee
}

func main() {
	//TODO ...
	//open test.txt and read, use defer to close the file
	//hint os package can be used to open the file, jsonbuilder package has functions
	//be sure to process errors

	//add a new employee to the list  hint:Employee{"ghi", 12000}

	//write the new list of employess to a file
	//again use jsonbuilder to write json formatted data
}
