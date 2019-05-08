package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
ai*(j-1)+bi*(n-j)

when ai > bi:
(ai-bi)*(j-1) + bi*(n-1)
sort by (ai-bi) desc

when bi > ai:
ai*(n-1) + (bi-ai)*(n-j)
sort by (bi-ai) asc

when bi == ai:
ai*(n-1)
in th middle

3
4 2
2 3
6 1

12

should optimize memory usage
*/

type Pair struct {
	a int64
	b int64
}

type Queue []Pair

func (q Queue) Score() (score int64) {
	n := int64(len(q))
	for j, p := range q {
		jj := int64(j + 1)
		score += p.a*(jj-1) + p.b*(n-jj)
	}

	return
}

func ScanInt(s *bufio.Scanner) (int, error) {
	s.Scan()
	return strconv.Atoi(s.Text())
}

func ScanInt64(s *bufio.Scanner) (int64, error) {
	s.Scan()
	return strconv.ParseInt(s.Text(), 10, 64)
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)
	sdtNum, _ := ScanInt(s)
	q1 := make([]Pair, 0)
	q2 := make([]Pair, 0)
	q3 := make([]Pair, 0)

	for i := 0; i < sdtNum; i++ {
		a, _ := ScanInt64(s)
		b, _ := ScanInt64(s)

		if a > b {
			q1 = append(q1, Pair{a, b})
		} else if a < b {
			q2 = append(q2, Pair{a, b})
		} else {
			q3 = append(q3, Pair{a, b})
		}
	}
	// ai > bi
	sort.Slice(q1, func(i, j int) bool {
		return (q1[i].a - q1[i].b) > (q1[j].a - q1[j].b)
	})
	// bi > ai
	sort.Slice(q2, func(i, j int) bool {
		return (q2[i].b - q2[i].a) < (q2[j].b - q2[j].a)
	})

	q := Queue(append(append(q1, q3...), q2...))

	fmt.Println(q.Score())
}
