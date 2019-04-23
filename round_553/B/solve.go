package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

The first line contains two integers 𝑛 and 𝑚 (1≤𝑛,𝑚≤500) — the number of rows and the number of columns in the matrix 𝑎.

inputCopy
3 2
0 0
0 0
0 0
outputCopy
NIE

inputCopy
2 3
7 7 7
7 7 10
outputCopy
TAK
1 3

1 2 3
4 5 6
7 8 9

*/

type Solver struct {
	a     [][]int
	n     int
	m     int
	Path  []int
	Found bool
}

func (s *Solver) pushPath(i int) {
	if s.Found == true {
		return
	}

	s.Path = append(s.Path, i)
}

func (s *Solver) popPath() (p int) {
	if s.Found == true {
		return
	}

	p, s.Path = s.Path[len(s.Path)-1], s.Path[:len(s.Path)-1]
	return
}

func (s *Solver) Result() {
	if !s.Found {
		fmt.Println("NIE")
		return
	}
	fmt.Println("TAK")
	fmt.Println(strings.Trim(fmt.Sprint(s.Path), "[]"))
	return
}

func (s *Solver) Search(res, depth int) {
	searched := map[int]bool{}

	for i := 0; i < s.m; i++ {
		if s.Found == true {
			break
		}

		cur_var := s.a[depth][i]

		// check if searched
		if searched[cur_var] {
			continue
		}
		searched[cur_var] = true

		s.pushPath(i + 1)
		if depth == s.n-1 {
			s.Found = (cur_var^res > 0)
		} else {
			s.Search(cur_var^res, depth+1)
		}
		s.popPath()
	}
}

func ScanNum(s *bufio.Scanner) (int, error) {
	s.Scan()
	return strconv.Atoi(s.Text())
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))

	s.Split(bufio.ScanWords)
	n, _ := ScanNum(s)
	m, _ := ScanNum(s)
	a := [][]int{}

	for i := 0; i < n; i++ {
		a = append(a, make([]int, m))
		for j := 0; j < m; j++ {
			a[i][j], _ = ScanNum(s)
		}
	}

	ss := Solver{
		a: a,
		n: n,
		m: m,
	}
	ss.Search(0, 0)
	ss.Result()
}
