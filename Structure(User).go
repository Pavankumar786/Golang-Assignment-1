9)Sturucture(user)
package main

import "fmt"

// defining a structure
type Employee struct {
	firstName, lastName   string
	age, dob,salary  int
	
}

func main() {

	
	emp:= &Employee{"Pavankumar", "Keerthi", 55,1997,6000}
        fmt.Println("First Name:", (*emp).firstName)
	fmt.Println("Last Name:", (*emp).lastName)
	fmt.Println("Age:", (*emp).age)
	fmt.Println("DOB:", (*emp).dob)
	fmt.Println("Salary:", (*emp).salary)
	
}

Output:
First Name: Pavankumar
Last Name: Keerthi
Age: 55
DOB: 1997
Salary: 6000