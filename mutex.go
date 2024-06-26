package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mutex sync.Mutex // synchronizacja dostepu do zmiennej counter pomiedzy gorutynami. Tylko jedna gorutyna może zablokować mutex w danym momencie. 

func increment() {
	mutex.Lock() // Zablokowanie mutexa, żeby tylko jedna gorutyna mogła mieć dostęp do zasobu.
	defer mutex.Unlock() // odblokowanie mutexa po zakończeniu funkcji increment. Defer zapewnia, że zostanie bezpośrednio odlobkowana po zakończeniu, niezależnie od tego czy kończy się normalnie czy z błędem. 

	// Blokowanie i odblokowywanie mutexa zapewnia, że zmiennej counter nie zmienią dwie gorutyny jednocześnie, co mogłoby prowadzić do nieprzewidywalnych wyników (takich jak błędne wartości counter).
	counter++
}

func decrement(){
	mutex.Lock()
	defer mutex.Unlock()
	counter--
}

func main() {
	var wg sync.WaitGroup // WaitGroup aby główny wątek 'main' czekał na zakończenie obu gorutyn, przed wpisaniem wartości counter. 
	wg.Add(2) // Dodajemy dwa ponieważ mamy dwie gorutyny. 

	go func() {
		defer wg.Done() // bez wg.Done() nigdy nie zakończy się działanie wg.wait(). Dekrementacja 
		for i := 0; i < 1000; i++ {
			increment()
		}
	}()

	go func() {
		defer wg.Done() 
		for i := 0; i < 1000; i++ {
			decrement()
		}
	}()
	// Dzięki temu, że każda z tych operacji jest chroniona przez mutex, wartość 'counter' nie będzie modyfikowana jednocześnie przez obie gorutyny. Obie funkcje będą wykonywane w sposób sekwencyjny.
	// Gdy usuniemy z kodu Lock i unlock dochodzi do datarace i wynik jest nieprzewidywalny (dwie gorutyny mogą równocześnie modyfikować wartość counter)

	wg.Wait()

	fmt.Println("Wartość licznika po operacjach:", counter)
	
}
