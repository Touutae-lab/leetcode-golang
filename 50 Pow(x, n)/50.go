package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		return pow(x, n)
	} else {
		return 1 / pow(x, -n)
	}
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	// Avoid Edge Case: n = -2147483648
	if n%2 == 0 {
		return pow(x*x, n/2)
		// Also avoid Edge Case: n = -2147483647
	} else {
		return x * pow(x*x, n/2)
	}
}
