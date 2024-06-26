package main

import (
	"fmt"
	"time"
)

func sum(nums []int, ch chan int) {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    ch <- sum
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //Tworzy slice (dynamiczną tablicę) nie trzeba określać jej rozmiaru
    ch := make(chan int, 4) // używamy buforowanego kanału dla 4 gorutyn (czyli się nie przyblokuje na jednej)

	start := time.Now() // zapisujemy czas rozpoczęcia pomiaru. 

    // Obliczanie wielkości części dla jednej gorutyny  
    partSize := len(nums) / 4

    // Uruchamianie 4 gorutyn, każda sumuje inną część tablicy
    for i := 0; i < 4; i++ {
        startIdx := i * partSize  
        endIdx := startIdx + partSize 
        if i == 3 { // ostatnia część może być większa, jeśli rozmiar nie dzieli się równo
            endIdx = len(nums)
        }
        go sum(nums[startIdx:endIdx], ch) // uruchamiamy gorutynę, która wywołuje funkcje sum dla fragmentu tablicy nums od indeksu start do end. 
    }

    // Odbieranie wyników z gorutyn
    total := 0
    for i := 0; i < 4; i++ {
        total += <-ch
    }

	elapsed := time.Since(start) // Obliczamy czas jaki minął od startu 

    fmt.Println("Total sum:", total)
	fmt.Println("Elapsed time", elapsed)
}

// Wytłumaczenie: 
// partSize := len(nums) / 4 = 10 / 4 = 2

// Dla i == 0:
// start = 0 * 2 = 0
// end = 0 + 2 = 2
// Gorutyna go sum(nums[0:2], ch) będzie sumować elementy 1 i 2.

// Dla i == 1:
// start = 1 * 2 = 2
// end = 2 + 2 = 4
// Gorutyna go sum(nums[2:4], ch) będzie sumować elementy 3 i 4.

// Dla i == 2:
// start = 2 * 2 = 4
// end = 4 + 2 = 6
// Gorutyna go sum(nums[4:6], ch) będzie sumować elementy 5 i 6.
// Dla i == 3:
// start = 3 * 2 = 6
// end = len(nums) = 10
// Gorutyna go sum(nums[6:10], ch) będzie sumować elementy 7, 8, 9 i 10.