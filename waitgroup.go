package main

import (
	"fmt"
	"sync"
	"time"
)

// stosujemy wskaźnik do sync.WaitGroup, bo chcemy aby każda gorutyna miała dostęp do tej samej instancji WaitGroup.
func worker(id int, start chan struct{}, wg *sync.WaitGroup ) {
	defer wg.Done() // Co każde wykonanie funkcji kończy swoją pracę, zmniejsza liczbę o 1.  
	fmt.Printf("Rozpoczęcie pracy pracownik %d\n", id)
	<-start // Pozwala wszystkim gorutynom zacząć pracę jednocześnie, gdy zostanie wysłany sygnał przez zamknięcie kanału.
	time.Sleep(time.Second) // zatrzymanie na sekunde, gorutyna czeka. 
	fmt.Printf("Zakończenie pracy pracownik %d\n", id)
} 

func main() {
	var wg sync.WaitGroup // struktura w go używana do synchronizacji gorutyn. 
	start := make(chan struct{}) //Deklaracja kanału start  

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Co każdą iteracje nowa gorutyna zaczyna swoją pracę.
		go worker(i, start, &wg) // Przekazanie 
	}

	close(start) // zakończenie kanału 

	wg.Wait() // czeka, aż wszystkie gorutyny zakończą działanie. 
	fmt.Println("Wszyscy pracownicy zakończyli pracę")
}
