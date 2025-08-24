package main

import "fmt"

type Employee interface {
	calculateSalary() int
}

type Fulltime struct {
	dailyRate  int
	daysWorked int
}
type Contractor struct {
	dailyRate  int
	daysWorked int
}
type Freelancer struct {
	hourlyRate  int
	hoursWorked int
}

func (f Fulltime) calculateSalary() int {
	return f.daysWorked * f.dailyRate
}
func (c Contractor) calculateSalary() int {
	return c.dailyRate * c.daysWorked
}
func (f Freelancer) calculateSalary() int {
	return f.hoursWorked * f.hourlyRate
}

func main() {
	fullTime := Fulltime{500, 30}
	contractor := Contractor{100, 30}
	freelancer := Freelancer{100, 20}

	fmt.Println(fullTime.calculateSalary())
	fmt.Println(contractor.calculateSalary())
	fmt.Println(freelancer.calculateSalary())

}
