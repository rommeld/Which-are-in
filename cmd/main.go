package main

import (
	"fmt"
	"whicharein"
)

func main() {
	arrayOne := []string{"tarp", "mice", "bull"}
	arrayTwo := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	example := whicharein.InArray(arrayOne, arrayTwo)
	fmt.Println(example)
}
