package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // Synchronizacja, główna gorutyna poczeka na zakończenie obu gorutyn. 
	wg.Add(2)
	// Utworzenie dwóch kanałów ch1 i ch2   
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Utworzenie dwóch go rutyn  
	go func() {
		defer wg.Done() // defer do odłożenia wykonania funkcji do momentu, gdy funkcja otaczająca zakończy swoje działanie. Ma to na celu zapewnienie, że wg.Done() zostanie wywołane po zakończeniu funkcji. 
		time.Sleep(2 * time.Second) // Po dwóch sekundach
		ch1 <- "Dane z ch1"
	}()

	go func() {
		defer wg.Done() // Decrementuje licznik gorutyn o 1. Jest odpowiednikiem Add(-1).
		time.Sleep(1 * time.Second) // Po jednej sekundzie
		ch2 <- "Dane z ch2"
	}()

	// Go rutyna która czeka na zakończenie obu go rutyn i zamyka kanały. 
	go func() {
		wg.Wait() // Blokuje wykonanie, aż licznik gorutyn osiągnie zero.
		close(ch1)
		close(ch2)
	}()

	for {
		select {
		case msg1, ok := <-ch1:
			if ok {
				fmt.Println("Odebrano z ch1:", msg1)
			} else {
				ch1 = nil // Ustawienie kanału na nil po zamknięciu. nil == null. 
			}
		case msg2, ok := <-ch2:
			if ok {
				fmt.Println("Odebrano z ch2:", msg2)
			} else {
				ch2 = nil // Ustawienie kanału na nil po zamknięciu
			}
		}

		// Wyjście z pętli, gdy oba kanały są zamknięte
		if ch1 == nil && ch2 == nil {
			break
		}
	}
}
