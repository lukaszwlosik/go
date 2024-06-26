package main

import (
	"fmt"
	"time"
)

// Funkcja sumująca liczby w podanym zakresie
func sum(nums []int, result chan int) {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    result <- sum // wysyłanie wyniku do kanału
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    result := make(chan int)

    start := time.Now()
    // Podzielenie pracy na dwie gorutyny
    go sum(nums[:len(nums)/2], result)
    go sum(nums[len(nums)/2:], result)

    // Odbieranie wyników z gorutyn
    sum1 := <-result
    sum2 := <-result

    total := sum1 + sum2

    elapsed := time.Since(start)
    fmt.Println("Total sum:", total)
    fmt.Println("Elapsed time:", elapsed)
}
