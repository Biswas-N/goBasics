package main

// multiple packages are imported like this
import (
	"fmt"
	"math"

	// When importing, import using full path starting from the GOPATH
	"github.com/biswas-n/goBasics/03_packages/mypackage"
)

func main() {
	fmt.Println("Packages")

	value := math.Ceil(10.33)
	fmt.Println(value)

	fmt.Println(mypackage.SayHi("Biswas"))
}
