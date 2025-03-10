package main

import (
	"fmt"
	"sync"
)

func countLetters(s string, ch chan map[rune]int, wg *sync.WaitGroup) {
    defer wg.Done()
    frequency := make(map[rune]int)
    for _, char := range s {
        if char >= 'a' && char <= 'z' {
            frequency[char]++
        }
    }
    ch <- frequency
}

func mergeFrequencies(ch chan map[rune]int) map[rune]int {
    finalFrequency := make(map[rune]int)
    for freq := range ch {
        for char, count := range freq {
            finalFrequency[char] += count
        }
    }
    return finalFrequency
}

func main() {
    strings := []string{"quick", "brown", "fox", "lazy", "dog"}
    ch := make(chan map[rune]int, len(strings))
    var wg sync.WaitGroup

    for _, s := range strings {
        wg.Add(1)
        go countLetters(s, ch, &wg)
    }

    wg.Wait()
    close(ch)

    finalFrequency := mergeFrequencies(ch)

    for char, count := range finalFrequency {
        fmt.Printf("%c: %d\n", char, count)
    }
}
