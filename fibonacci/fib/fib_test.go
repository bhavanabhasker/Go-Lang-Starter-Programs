package fib

import "testing"

func Test(t *testing.T) {
	var v uint
	v = fibo(10)
	if v != 55 {
		t.Error(" got", v)
	}
}
