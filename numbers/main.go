package main

	
import (
    "fmt"
    "sync"
)

func printEven(n int, wg *sync.WaitGroup, chin *chan int, chout *chan int) {
	// fmt.Printf("Start printEven %d \n", n)
	// fmt.Printf("Start printEven %d %d\n", (*chin), (*chout))
	<-(*chin)
	
		if n % 2 == 0 {
			fmt.Printf("Even %d \n", n)
		}
	(*chout)<-n
	wg.Done()
}

func printOdd(n int, wg *sync.WaitGroup, chin *chan int, chout *chan int) {
	<-(*chin)
		if n % 2 !=  0 {
			fmt.Printf("Odd %d \n", n)
		}
	(*chout)<-n
	wg.Done()
}

func main() {
	const N = 20
	var wg sync.WaitGroup

	var chin = []chan int{}
	var chout = []chan int{}

    for i := 0; i < N+1; i++ {
		chin = append(chin, make(chan int))
        chout = append(chout, make(chan int))
    }

	wg.Add(2*N)
	fmt.Println("Init")
	
	for i:=0 ; i < N; i++ {
		go printEven(i, &wg, &(chin[i]), &(chout[i]))
		go printOdd(i, &wg, &chout[i], &chin[i+1])
	} 
	chin[0]<- -1
	<-chin[N]
	wg.Wait()
}