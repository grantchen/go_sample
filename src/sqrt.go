package main

import "fmt"
import "math"

func sqrt(x float64) float64 {
	x_old := 1.0
	x_new := 0.0
	for {
		x_new = (x_old + x/x_old) / 2
		if math.Abs(x_new-x_old) < 0.000001 {
			break
		}
		x_old = x_new
	}
	return x_new
}

func main() {
	fmt.Println(sqrt(2))
}
