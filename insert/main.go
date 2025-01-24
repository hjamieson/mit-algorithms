package insert

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("main")
	sorted := Sort([]int{5, 4, 3, 2, 1})
	fmt.Println(sorted)
	os.Exit(0)
}
