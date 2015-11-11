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
	f, err := os.Open("test.txt")
	defer f.Close()
	emprecs := new(EmpRecs)
	err = jsonbuilder.ReadData(f, &emprecs.emps)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(emprecs)

	emprecs.emps = append(emprecs.emps, Employee{"ghi", 12000})

	f2, err := os.Create("test2.txt")
	defer f2.Close()
	jsonbuilder.WriteData(f2, emprecs.emps)
	fmt.Println(emprecs)
}
