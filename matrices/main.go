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

func addSimple(a, b, r *int, wg *sync.WaitGroup) {
	*r = *a + *b
	(*wg).Done()
}

func main() {
	const N, M = 5, 3
	m1 := make([][]int, N)
	m2 := make([][]int, N)
	result := make([][]int, N)

	buildMatrix(&m1, N, M)
	buildMatrix(&m2, N, M)

	fmt.Printf("Start calculation\n")

	for i := 0; i < N; i++ {
		result[i] = make([]int, M)
	}

	var wg sync.WaitGroup
	wg.Add(N * M)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			go addSimple(&m1[i][j], &m2[i][j], &result[i][j], &wg)
		}
	}

	wg.Wait()

	fmt.Printf("%v\n", m1)
	fmt.Printf("%v\n", m2)

	fmt.Printf("%v\n", result)
}
