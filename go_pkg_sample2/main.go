package main

import (
	"employee"
	"fmt"
	"jsonbuilder"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	defer f.Close()
	emprecs := new(employee.EmpRecs)
	err = jsonbuilder.ReadData(f, &emprecs.Emps)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(emprecs)

	emprecs.Emps = append(emprecs.Emps, employee.Employee{"ghi", 12000, "mgr1"})
	f2, err := os.Create("test2.txt")
	defer f2.Close()
	jsonbuilder.WriteData(f2, emprecs.Emps)
	fmt.Println(emprecs)
}
