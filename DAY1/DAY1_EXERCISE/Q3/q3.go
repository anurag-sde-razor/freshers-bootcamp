package main

import (
    "fmt"
)

type Employee interface {
    CalculateSalary() int
}


type FullTime struct {
    SalaryPerMonth int
}


type Contractor struct {
    SalaryPerMonth int
}


type Freelancer struct {
    RatePerHour int
    HoursWorked int
}

func (f FullTime) CalculateSalary() int {
    return f.SalaryPerMonth
}

func (c Contractor) CalculateSalary() int {
    return c.SalaryPerMonth
}


func (fr Freelancer) CalculateSalary() int {
    return fr.RatePerHour * fr.HoursWorked
}

func main() {
    
    fullTimeEmployee := FullTime{SalaryPerMonth: 15000}
    contractorEmployee := Contractor{SalaryPerMonth: 3000}
    freelancerEmployee := Freelancer{RatePerHour: 100, HoursWorked: 20}

    employees := []Employee{fullTimeEmployee, contractorEmployee, freelancerEmployee}

    
    for _, employee := range employees {
        fmt.Printf("Salary: %d\n", employee.CalculateSalary())
    }
}