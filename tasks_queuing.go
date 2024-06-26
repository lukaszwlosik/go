// Program służy jako system kolejkowania zadań
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Utworzenie struktury dla zadania
type Task struct {
	ID       int // każde zadanie ma swoje id
	Duration time.Duration // każde zadanie ma swój czas wykonania 
}
// id - pracownika, tasks - kanał do odczytu, results - kanał do zapisu, wg - wskaźnik do sync.WaitGroup, synchronizacja gorutyn. 
func Worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // dekrementacja licznika WaitGroup
	for task := range tasks {
		time.Sleep(task.Duration) // Simulate task processing.
		results <- fmt.Sprintf("Pracownik %d wykonuje zadanie nr: %d", id, task.ID) // wysłanie stringa do kanału
	}
}

func main() {
	numWorkers := 5   
	numTasks := 20
	// kanał buforowany - określona pojemność. Możemy określić liczbę wartości, zanim konieczne będzie ich odebranie przez inną gorutynę (Nie blokują)
	tasks := make(chan Task, numTasks) // utowrzenie kanału buforowanego tasks
	results := make(chan string, numTasks) // utworzenie kanału buforowanego results 
	var wg sync.WaitGroup //synchronizacja gorutyn 

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, tasks, results, &wg) // każdy pracownik wykonuje funkcje i zwiększa waitgroup o 1. 
	}

	for i := 1; i <= numTasks; i++ {
		tasks <- Task{
			ID:       i,
			Duration: time.Duration(rand.Intn(1000)) * time.Millisecond,
			// wysłanie do kanału tasks zadań z losowym czasem wykonania.
		}
	}
	close(tasks) // zamknięcie kanału tasks

	go func() {
		wg.Wait() // Czeka aż wszystkie gorutyny zakończą działanie żeby zamknąć kanał.
		close(results) // zamknięcie kanału results
	}()

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("Wszystkie zadania zostały wykonane pomyślnie.")
}
