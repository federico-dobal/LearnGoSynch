package main

import (
	"fmt"
	"sync"
)

func buildMatrix(M *[][]float64, n, m int) {
	c := make([]float64, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c[j] = float64(j + i)
		}
		(*M)[i] = make([]float64, m)
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

func addSimple(m [][]float64, i, j int, wg *sync.WaitGroup, mu *sync.Mutex) {
	(*mu).Lock()
	var avg = 0.0
	var un = 1
	if i > 0 {
		avg += m[i-1][j]
		un++

		if j > 0 {
			avg += m[i-1][j-1]
			un++
		}

		if j < len(m[0])-1 {
			avg += m[i-1][j+1]
			un++
		}
	}

	if j > 0 {
		avg += m[i][j-1]
		un++
	}

	if j < len(m[0])-1 {
		avg += m[i][j+1]
		un++
	}

	if i < len(m)-1 {
		avg += m[i+1][j]
		un++

		if j > 0 {
			avg += m[i+1][j-1]
			un++
		}

		if j < len(m[0])-1 {
			avg += m[i+1][j+1]
			un++
		}
	}

	m[i][j] = avg / float64(un)
	(*mu).Unlock()
	(*wg).Done()
}

func main() {
	const N, M = 5, 3
	m1 := make([][]float64, N)

	buildMatrix(&m1, N, M)

	fmt.Printf("Start calculation\n")

	var wg sync.WaitGroup
	var m sync.Mutex

	fmt.Printf("%v\n", m1)

	wg.Add(N * M)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			go addSimple(m1, i, j, &wg, &m)
		}
	}

	wg.Wait()

	fmt.Printf("%v\n", m1)
}
