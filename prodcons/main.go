package main

	
import (
    "fmt"
	"sync"
	"math/rand"
	"time"
)

func mainUpdate(v *int, wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1000)
	fmt.Printf("Update %d\n", n)
	*v = n
	(*wg).Done()
}

func procUpdate(v *int, pv *int, wg *sync.WaitGroup) {
	*pv = *v
	(*wg).Done()
}

func main() {
	var v, pv1, pv2 int 
	v, pv1, pv2 = 10, 5, 7
	var wg sync.WaitGroup


	for i:=0 ; i < 10 ; i++ {
		fmt.Printf("Before update %d, %d, %d\n", v, pv1, pv2)
		wg.Add(1)
		go mainUpdate(&v, &wg) 
		wg.Wait()
		wg.Add(2)
		go procUpdate(&v, &pv1, &wg) 
		go procUpdate(&v, &pv2, &wg) 
		wg.Wait()
		fmt.Printf("After update %d, %d, %d\n", v, pv1, pv2)
	}
}