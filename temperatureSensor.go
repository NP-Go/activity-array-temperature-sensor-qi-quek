package main

import "fmt"

func main() {
	temperature := [24]float64{20.1, 24, 27.3, 30.1, 26.4, 22.2, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3}

	var total float64

	for i := 0; i < len(temperature); i++ {
		total += temperature[i]
	}
	fmt.Printf("Average temperature of the day is %.02f", (total / 24))
}
