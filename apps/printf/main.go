package main

import (
	"fmt"

	"github.com/chaitsgithub/hfpackages/packages/greeting"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	counters := [3]int{5, 20, 17}
	fmt.Println(counters)
	fmt.Printf("%#v\n", counters)
}
