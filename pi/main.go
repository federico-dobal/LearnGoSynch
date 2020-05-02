package main

import (
	"fmt"
	"sync"
)

func buildMatrix(M *[][]int, n, m int) {
	c := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c[j] = j + i
		}
		(*M)[i] = make([]int, m)
		copy((*M)[i], c[:])
	}
}

func add(m1, m2, result *[][]int) {
	for i := 0; i < len(*m1); i++ {
		c := len((*m1)[i])
		for j := 0; j < c; j++ {
			((*result)[i])[j] = (*m1)[i][j] + (*m2)[i][j]
		}
	}
}

func calculatePiStep(i int, step float64, x, sum *float64, m *sync.Mutex, wg *sync.WaitGroup) {
	(*m).Lock()
	*x = (float64(i) + 0.5) * step
	*sum = (*sum) + 4.0/(1.0+(*x)*(*x))
	(*m).Unlock()
	(*wg).Done()
}

func calculatePiGo(numSteps int) {
	x := 0.0
	sum := 0.0
	step := 1.0 / float64(numSteps)

	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(numSteps)

	for i := 0; i < numSteps; i++ {
		go calculatePiStep(i, step, &x, &sum, &m, &wg)
	}

	wg.Wait()
	pi := step * sum
	fmt.Printf("pi go = %f\n", pi)
}

func calculatePi(numSteps int) {
	x := 0.0
	sum := 0.0
	step := 1.0 / float64(numSteps)
	for i := 0; i < numSteps; i++ {
		x = (float64(i) + 0.5) * step
		sum = sum + 4.0/(1.0+x*x)
	}

	pi := step * sum
	fmt.Printf("pi = %f\n", pi)
}

func main() {
	const numSteps = 1000
	calculatePiGo(numSteps)
	calculatePi(numSteps)
}
