package main

import (
	"fmt"

	"github.com/chaitsgithub/hfpackages/packages/magazine"
)

func main() {
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
}
