package main

	
import (
    "fmt"
	"sync"
	"time"
)

func phil(id int, wg *sync.WaitGroup, fl *sync.Mutex, fr *sync.Mutex) {
	fl.Lock()
	time.Sleep(2 * time.Second)
	fr.Lock()
	fmt.Printf("Eating %d\n"	, id)
	fl.Unlock()
	fr.Unlock()

	fmt.Printf("Thinking %d\n", id)
	
	(*wg).Done()
}

func main() {
	const N = 20
	var wg sync.WaitGroup
	var forks []sync.Mutex

	for i:=0 ; i < N+1 ; i++ {	
		forks = append(forks, sync.Mutex{})
	}

	wg.Add(N)
	for i:=0 ; i < N ; i++ {
		fmt.Printf("Mutex len: %d - %d\n", i, (i+1) % N)			
		// Condition to avoid deadlock if all the philosophers decides to take firs left
		// fork
		if i % 2 == 0 {
			// First pick left fork
			go phil(i, &wg, &forks[i], &forks[(i+1) % N]) 
		} else {
			// First pick right fork
			go phil(i, &wg, &forks[(i+1) % N], &forks[i]) 
		}
		
	}
	wg.Wait()
}