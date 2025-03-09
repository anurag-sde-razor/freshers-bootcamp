package main

import (
    "fmt"
    "sync"
)

type BankAccount struct {
    balance int
    mu      sync.Mutex
}

func (a *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
    defer wg.Done()
    a.mu.Lock()
    defer a.mu.Unlock()
    a.balance += amount
    fmt.Printf("Deposited %d, new balance: %d\n", amount, a.balance)
}

func (a *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
    defer wg.Done()
    a.mu.Lock()
    defer a.mu.Unlock()
    if a.balance >= amount {
        a.balance -= amount
        fmt.Printf("Withdrew %d, new balance: %d\n", amount, a.balance)
    } else {
        fmt.Printf("Failed to withdraw %d, balance: %d\n", amount, a.balance)
    }
}

func main() {
    account := &BankAccount{balance: 500}
    var wg sync.WaitGroup

    // Simulate concurrent deposits and withdrawals
    transactions := []struct {
        amount int
        isDeposit bool
    }{
        {100, true},
        {50, false},
        {200, true},
        {300, false},
        {150, true},
        {100, false},
    }

    for _, t := range transactions {
        wg.Add(1)
        if t.isDeposit {
            go account.Deposit(t.amount, &wg)
        } else {
            go account.Withdraw(t.amount, &wg)
        }
    }

    wg.Wait()
    fmt.Printf("Final balance: %d\n", account.balance)
}