package main

import (
	"flag"
	"fmt"
)

type KiloMeter float64
type Meter float64

const ONE_KM KiloMeter = 1

func KMToMeter(km KiloMeter) (m Meter) {
	m = Meter(km * 1000)
	return
}

func MeterToKiloMeter(m Meter) (km KiloMeter) {
	km = KiloMeter(m / 1000)
	return
}

var b = flag.Bool("n", false, "this is a boolean flag value, you see, right?")

func main() {
	fmt.Println("main invoked")

	flag.Parse()
	fmt.Println("value of b flag is: ", *b)

	a := ONE_KM
	var b Meter = 1000

	// fmt.Println(a == b)
	fmt.Println(a == 1)
	fmt.Println(a == KiloMeter(b))
	fmt.Println(a == 1)
}
