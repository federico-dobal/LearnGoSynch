package main

	
import (
    "fmt"
	// "sync"
)

func update(v *int) {
	if (*v) == 2 {
		(*v) = 0
	} else {
		*v = (*v) + 1	
	}
}

func main() {
	const N = 10
	var s1 int 
	var s2 int 
	
	s1 = 0
	s1 = 2

	// wg.Add(N)
	for i:=0 ; i < N ; i++ {
		fmt.Printf("Before update %d, %d\n", s1, s2)
		
		update(&s1) 
		update(&s2) 
		
		fmt.Printf("After update %d, %d\n", s1, s2)
	}
	
	// wg.Wait()
	// fmt.Printf("End update %d\n", v)
}