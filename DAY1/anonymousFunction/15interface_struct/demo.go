package main

import "fmt"

type payment interface {
	pay() string
}

type CreditCard struct {
	CardNumb string
	Amount   int
}

type Upi struct {
	Upi_id string
	Amount int
}

type CashOnDelivery struct {
	Amount int
}

func (c CreditCard) pay() string {
	return fmt.Sprintf("\nYour payment Processing Through Credit card method:\n\n Amount: %d\n Card Number: %s\n", c.Amount, c.CardNumb)
}
func (u Upi) pay() string {
	return fmt.Sprintf("\nYour payment Processing Through Upi method:,\n Amount:%d\n Upi Id:%s \n", u.Amount, u.Upi_id)
}

func (cod CashOnDelivery) pay() string{
	return fmt.Sprintf("\nYour payment Processing Through Cash on Delivery method:\n\n Amount:%d\n", cod.Amount)
}

func main() {
	fmt.Println("Enter your First name")
	str := ""

	fmt.Scanln(&str)

	fmt.Println("\nWelcome ", str, "!")

	fmt.Println("\nChoose Which Payment Method You Want To Use:\n")
	fmt.Println("1.Credit Card")
	fmt.Println("2.Upi")
	fmt.Println("3.Cash On Delivery")
	fmt.Println("\nEnter Your Choice:")
	choice := 0

	fmt.Scan(&choice)
	    if choice > 3 || choice < 1 {
	    fmt.Println("Invalid Input!")
		return	
	}

		cardnumber := ""
		amount1 := 0
		if choice == 1 {
			fmt.Println("\nEnter your Card Number:")
			fmt.Scan(&cardnumber)
			fmt.Println("\nEnter Amount:")
			fmt.Scan(&amount1)

		}
		fmt.Println("\nEnter your Card Number:")
		upi1 := ""
		amount2 := 0

		if choice == 2 {
			fmt.Println("\nEnter your Upi Id:")
			fmt.Scan(&upi1)
			fmt.Println("\nEnter Amount:")
			fmt.Scan(&amount2)

		}

		amount3 := 0
		if choice == 3 {

			fmt.Println("\nEnter Amount:")
			fmt.Scan(&amount3)

		}

		CreditCard1 := CreditCard{
			CardNumb: cardnumber,
			Amount:   amount1,
		}
		Upi1 := Upi{
			Upi_id: upi1,
			Amount: amount2,
		}
		CashOnDelivery1 := CashOnDelivery{
			Amount: amount3,
		}

		payment1 := []payment{
			CreditCard1,
			Upi1,
			CashOnDelivery1,
		}

		for index, value := range payment1 {
			if index+1 == choice {
				fmt.Println(value.pay())
				break
			}
		}

	}

