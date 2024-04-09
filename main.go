package main

import (
	"fmt"
	"github.com/Mexx77/golang-mono/pkg/generic_data_types/set"
)

func main() {
	myInt := 42
	intSet := set.Make[int]()
	intSet.Add(myInt)
	fmt.Printf("checking if %d is contained in the set: %t\n", myInt, intSet.Contains(myInt))
}
