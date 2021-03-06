// UVa 530 - Binomial Showdown

package main

import (
	"fmt"
	"math/big"
	"os"
)

func f(n int) big.Int {
	var r, t big.Int
	r.SetInt64(1)
	for i := 2; i <= n; i++ {
		t.SetInt64(int64(i))
		r.Mul(&r, &t)
	}
	return r
}

func c(n, m int) big.Int {
	r := f(n)
	s := f(n - m)
	t := f(m)
	r.Div(&r, &s)
	r.Div(&r, &t)
	return r
}

func main() {
	in, _ := os.Open("530.in")
	defer in.Close()
	out, _ := os.Create("530.out")
	defer out.Close()

	var n, m int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		r := c(n, m)
		fmt.Fprintln(out, &r)
	}
}
