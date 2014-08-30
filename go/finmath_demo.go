package main

import (
	"fmt"

	"github.com/vagababov/finmath/go/finmath"
)

func main() {
	fmt.Println(finmath.PMT(350e3, 12*30, 0.03625/12))
}
