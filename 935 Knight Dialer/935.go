package main

func knightDialer(n int) int {
	calfour := (3^n-1)%1e9 + 7
	calother := (9*(2^n-1))%1e9 + 7
	calfive := n
	return (calfour + calother + calfive) % 1e9

}
