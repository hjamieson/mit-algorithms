package main

import (
	"fmt"
	"mit/algorithms/insert"
	"os"
)

func main() {
	fmt.Println("main")
	sorted := insert.Sort([]int{5, 4, 3, 2, 1})
	fmt.Println(sorted)
	os.Exit(0)
}
