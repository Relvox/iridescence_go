package main

import (
	"fmt"

	"github.com/relvox/iridescence_go/utils"
)

type dummy struct {
	name string
	num  int
}

func main() {
	var x0 int = 5
	var x1 uint = 0
	var x2 float64 = 0
	var x3 dummy = dummy{"A", 1}
	var x4 *dummy = &dummy{"A", 1}
	var x5 *dummy = nil
	var x6 any = dummy{"A", 1}
	var x7 any = &dummy{"A", 1}
	var x8 any = nil

	fmt.Printf("0. %v %x %t\n", x0, x0, utils.IsNilOrZero(x0))
	fmt.Printf("1. %v %x %t\n", x1, x1, utils.IsNilOrZero(x1))
	fmt.Printf("2. %v %x %t\n", x2, x2, utils.IsNilOrZero(x2))
	fmt.Printf("3. %v %x %t\n", x3, x3, utils.IsNilOrZero(x3))
	fmt.Printf("4. %v %x %t\n", x4, x4, utils.IsNilOrZero(x4))
	fmt.Printf("5. %v %x %t\n", x5, x5, utils.IsNilOrZero(x5))
	fmt.Printf("6. %v %x %t\n", x6, x6, utils.IsNilOrZero(x6))
	fmt.Printf("7. %v %x %t\n", x7, x7, utils.IsNilOrZero(x7))
	fmt.Printf("8. %v %x %t\n", x8, x8, utils.IsNilOrZero(x8))

}
