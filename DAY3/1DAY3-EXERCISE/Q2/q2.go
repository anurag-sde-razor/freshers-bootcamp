package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func getRating(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    rating := rand.Intn(10) + 1 
    ch <- rating
}

func main() {
    rand.Seed(time.Now().UnixNano())
    const numStudents = 200
    ch := make(chan int, numStudents)
    var wg sync.WaitGroup

    for i := 0; i < numStudents; i++ {
        wg.Add(1)
        go getRating(ch, &wg)
    }

    wg.Wait()
    close(ch)

    totalRating := 0
    count := 0
    for rating := range ch {
        totalRating += rating
        count++
    }

    averageRating := float64(totalRating) / float64(count)
    fmt.Printf("Average Rating: %.2f\n", averageRating)
}