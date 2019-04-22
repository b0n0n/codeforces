package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
target string ACTG

input:
4
ZCTH
output:
2
*/
func validUpperByte(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

func Distance(a, b byte) (int, error) {
	if !validUpperByte(a) || !validUpperByte(b) {
		return -1, fmt.Errorf("Invalid bytes")
	}

	var gap int

	if a > b {
		gap = int(a - b)
	} else {
		gap = int(b - a)
	}

	if gap > 13 && gap < 26 {
		return 26 - gap, nil
	}
	return gap, nil
}

func mutate(pattern, source []byte) (minOp int) {
	minOp = 9999

	p_l, s_l := len(pattern), len(source)

	if s_l < p_l {
		return minOp
	}

	for i := 0; i+p_l-1 <= s_l-1; i += p_l {
		tmpOp := 0
		for j := range pattern {
			dis, err := Distance(pattern[j], source[i+j])
			if err != nil {
				panic(err)
			}
			tmpOp += dis
		}
		if tmpOp < minOp {
			minOp = tmpOp
		}
	}

	return
}

func main() {
	var genome = []byte("ACTG")
	minOp := 9999

	// parse input
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	s.Scan() // discard len
	s.Scan()
	str := s.Bytes()

	for i := range genome {
		curOp := mutate(genome, str[i:])
		if curOp < minOp {
			minOp = curOp
		}
	}

	fmt.Println(minOp)
}
