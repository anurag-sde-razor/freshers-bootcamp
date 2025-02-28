package main

import "fmt"

type payment struct {
	amount int
	method string
}

func (p payment) pay() {
	fmt.Println("payment Amount:", p.amount, "\nMethod:", p.method)
}

func main() {
	p1 := payment{
		amount:100,
		method:"UPI",
	}

	p1.pay()

}
