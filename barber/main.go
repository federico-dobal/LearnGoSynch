package main

import (
	"fmt"
	"sync"
)

func barber(nfs int, br *sync.Mutex, cr *sync.Mutex, s *sync.Mutex) {
	(*cr).Lock()
	fmt.Printf("barber Seat lock\n")
	(*s).Lock()
	nfs++
	(*br).Unlock()
	(*s).Unlock()
	fmt.Printf("Barber cut hair!!!\n")
}

func customer(nfs int, br *sync.Mutex, cr *sync.Mutex, s *sync.Mutex) {
	fmt.Printf("customer Seat lock\n")
	(*s).Lock()
	if nfs > 0 {
		nfs--
		(*cr).Unlock()
		(*s).Unlock()
		(*br).Lock()
		fmt.Printf("Customer have hair cut!!!\n")
	} else {
		(*s).Unlock()
		fmt.Printf("Customer have NO hair cut!!! :(\n")
	}
}

func main() {
	const N = 20

	var br sync.Mutex
	var s sync.Mutex
	var cr sync.Mutex
	var nfs = N

	var wg sync.WaitGroup

	fmt.Printf("Seat lock\n")
	s.Lock()

	wg.Add(2)

	go barber(nfs, &br, &cr, &s)
	go customer(nfs, &br, &cr, &s)

	wg.Wait()
}
