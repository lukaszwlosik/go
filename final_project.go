// Program w którym każda gorutyna oblicza statystyki dla swojej części tablicy i wysyła do kanału result.
package main

import (
	"fmt"
	"math"
	"sync"
)

//Przekazanie wskaźnika  do początku tablicy, wraz z jej długością. Tablice przekazywane przez referencję. Funkcja zliczająca sumę.
func sum(nums []int) int {
	total := 0
	//Iteracja po elementach tablicy nums
	for _, num := range nums {
		total += num
	}
	return total
}
// Funkcja zliczająca średnią jako double. 
func average(nums []int) float64 {
	// Warunek dla zerowej tablicy. 
	if len(nums) == 0 {
		return 0
	}
	// Zwrócenie sumy i przypisanie do zmiennej
	sum := sum(nums)
	avg := float64(sum) / float64(len(nums))
	return avg
}

func min(nums []int) int {
	// Jeśli tablica nie istnieje przypisz 0 
	if len(nums) == 0 {
		return 0
	}
	min := math.MaxInt32
	for _, num := range nums {
		// Sprawdzenie minimalngo numeru 
		if num < min {
			min = num
		}
	}
	return min
}

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// Największa możliwa wartość dla 32-bitowej liczby całkowitej ze znakiem.
	max := math.MinInt32
	for _, num := range nums {
		// Warunek dla maksymalnej liczby 
		if num > max {
			max = num
		}
	}
	return max
}
// result chan<- map[string]interface{} - Result jest kanałem, do którego można tylko wysłać dane, ale nie można ich odbierać.  
func calculateStatistics(nums []int, wg *sync.WaitGroup, result chan<- map[string]interface{}) {
	defer wg.Done() // zmniejsza licznik o jeden gdy zakończy swoje działanie, przeciwieństwo add.

	// Utworzenie mapy, gdzie kluczami są napisy, a wartościami dowolne typy.
	stats := make(map[string]interface{})

	// Przypisanie wyników
	stats["Suma wynosi"] = sum(nums)
	stats["Średnia wynosi"] = average(nums)
	stats["Minimalna wynosi"] = min(nums)
	stats["Maksymalna wynosi"] = max(nums)
	
	// wysłanie wartości do kanału result  
	result <- stats
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65}

	// Podział listy na fragmenty
	numParts := 4
	partSize := len(nums) / numParts

	// Synchronizacja gorutyn. Pomaga w oczekiwaniu na zakończenie grupy gorutyn
	var wg sync.WaitGroup
	// Zwiększenie licznika gorutyn
	wg.Add(numParts)

	// Utworzenie kanału typu mapa, gdzie kluczami są napisy, a wartościami mogą być dowolne typy.
	// interface{} pozwala na przechowywanie w mapie wartości dowolnego typu. 
	result := make(chan map[string]interface{})

	// Uruchomienie gorutyn do obliczeń dla każdego fragmentu
	for i := 0; i < numParts; i++ {
		start := i * partSize
		end := start + partSize
		if i == numParts-1 { // warunek dla ostatnij częśći, może być większa
			end = len(nums)
		}
		// Wywołanie gorutyny dla tablicy dla poszczególnych fragmentów,oraz wskaźnika do sync.WaitGroup, aby wszystkie gorutyny operowały na tej samej instancji, oraz kanału. 
		go calculateStatistics(nums[start:end], &wg, result)
	}

	// partSize := len(nums) / numparts = 10 / 4 = 2

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

	// Utworzenie nowej go rutyny
	go func() {
		wg.Wait() // blokuje główną gorutynę, dopóki wszystkie gorutyny nie skończą pracy i nie wywołają wg.Done()
		close(result) // Zamykanie kanału, wszystkie dane zostały wysłane i przetworzone.
	}()

	// Przetwarzanie wartości przesłane przez kanał result, dopóki nie zostaną odczytane wszystkie wartości z kanału. 
	i := 0
	for stats := range result {
		i++
		fmt.Println("Wyniki dla fragmentu:", i, stats)
	}

	fmt.Println("Wszystkie obliczenia zostały zakończone.")
}
