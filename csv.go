package main

// START OMIT
import (
	_ "embed"
	"fmt"
)
//go:embed sample.csv
var csvData string

func main() {
	fmt.Print(csvData)
}
// END OMIT