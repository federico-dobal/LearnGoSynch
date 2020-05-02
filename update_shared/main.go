package main

	
import (
    "fmt"
	"sync"
)

func mainUpdate(v *int, wg *sync.WaitGroup, m *sync.Mutex) {
	(*m).Lock()
	*v = (*v) + 1
	(*m).Unlock()
	(*wg).Done()
}

func main() {
	const N = 100
	var v, pv1, pv2 int 
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(N)
	for i:=0 ; i < N ; i++ {
		fmt.Printf("Before update %d, %d, %d\n", v, pv1, pv2)
		
		go mainUpdate(&v, &wg, &m) 
		
		fmt.Printf("After update %d, %d, %d\n", v, pv1, pv2)
	}
	wg.Wait()
	fmt.Printf("End update %d\n", v)
}