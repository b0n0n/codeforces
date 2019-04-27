package main

/*

is there any painless way to use bignum lib??

*/

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

var N = big.NewInt(1000000007)
var Q = big.NewInt(4)
var D = big.NewInt(2)
var B0 = big.NewInt(0)
var B1 = big.NewInt(1)
var B2 = big.NewInt(2)
var B3 = big.NewInt(3)

// calculate the last number of the arithmetic progression
func c_li(f_i, n, d *big.Int) *big.Int {
	return new(big.Int).Add(f_i, new(big.Int).Mul(d, new(big.Int).Sub(n, B1)))
}

// calculate the first number of the arithmetic progression
func c_fi(l_i, n, d *big.Int) *big.Int {
	return new(big.Int).Sub(l_i, new(big.Int).Mul(d, new(big.Int).Sub(n, B1)))
}

// calculate the arithmetic progression sum
func APSum(f_i, l_i, d *big.Int) *big.Int {
	a := new(big.Int).Add(f_i, l_i)
	b := new(big.Int).Div(new(big.Int).Sub(l_i, f_i), d)
	b = new(big.Int).Add(b, B1)
	return new(big.Int).Div(new(big.Int).Mul(a, b), B2)
}

/*
x can be in section A or B
===============
        A   v   B
acccceeeeeeeeeeeeeeee
bbdddddddd
===============
*/
func seqSumEven(x *big.Int) (res *big.Int) {
	if x.Cmp(B1) == 0 {
		return B1
	}
	lx := big.NewInt(int64(math.Log2(float64(x.Int64()))))
	i2_u := new(big.Int).Sub(B2, new(big.Int).Mul(B2, new(big.Int).Exp(Q, lx.Div(lx, B2), nil)))
	i2_d := new(big.Int).Sub(B1, Q)
	i2 := new(big.Int).Mul(i2_u.Div(i2_u, i2_d), B2)
	if i2.Cmp(x) < 0 { // Section B
		n := new(big.Int).Sub(x, i2)
		i2a1 := new(big.Int).Add(i2, B1)
		l_i := c_li(i2a1, n, D)
		res = new(big.Int).Add(APSum(B1, i2, B1), APSum(i2a1, l_i, D))
	} else { // Section A
		n := new(big.Int).Sub(i2, x)
		i2s1 := new(big.Int).Sub(i2, B1)
		f_i := c_fi(i2s1, n, D)
		res = new(big.Int).Sub(APSum(B1, i2, B1), APSum(f_i, i2s1, D))
	}
	return
}

/*
x can be in section A or B
================
acccc
bbdddddddd
   A  ^ B
================
*/
func seqSumOdd(x *big.Int) (res *big.Int) {
	if x.Cmp(B2) == 0 {
		return B3
	}
	if x.Cmp(B3) == 0 {
		return big.NewInt(7)
	}

	lx := big.NewInt(int64(math.Log2(float64(x.Int64()))))
	i1_u := new(big.Int).Sub(B1, new(big.Int).Exp(Q, lx.Div(new(big.Int).Add(lx, B1), B2), nil))
	i1_d := new(big.Int).Sub(B1, Q)
	i1 := new(big.Int).Mul(i1_u.Div(i1_u, i1_d), B2)

	if i1.Cmp(x) < 0 { // Section B
		n := new(big.Int).Sub(x, i1)
		i1a2 := new(big.Int).Add(i1, B2)
		l_i := c_li(i1a2, n, D)
		res = new(big.Int).Add(APSum(B1, i1, B1), APSum(i1a2, l_i, D))
	} else { // Section A
		n := new(big.Int).Sub(i1, x)
		f_i := c_fi(i1, n, D)
		res = new(big.Int).Sub(APSum(B1, i1, B1), APSum(f_i, i1, D))
	}
	return
}

// SeqSum returns the sum of the sequence from the first element to the xth element
func SeqSum(x *big.Int) (res *big.Int) {
	if x.Cmp(B0) == 0 {
		return B0
	}
	lx := big.NewInt(int64(math.Log2(float64(x.Int64()))))
	if lx.Mod(lx, B2).Cmp(B0) == 0 {
		res = seqSumEven(x)
	} else {
		res = seqSumOdd(x)
	}

	return res
}

func solve(l, r int64) {
	rr := SeqSum(big.NewInt(r))
	lr := SeqSum(big.NewInt(l - 1))
	res := new(big.Int).Mod(rr.Sub(rr, lr), N)
	fmt.Println(res)
}

func ScanNum(s *bufio.Scanner) (int64, error) {
	s.Scan()
	// old version atoi default set to int32 instead of int64
	return strconv.ParseInt(s.Text(), 10, 64)
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	l, _ := ScanNum(s)
	r, _ := ScanNum(s)

	if l > r {
		panic(fmt.Errorf("Invalid range"))
	}

	solve(int64(l), int64(r))
}
