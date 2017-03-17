package main

import (
	"fmt"

	"github.com/naqernf/buildinfo"
)

func main() {
	fmt.Printf("Hallo World\n")
	fmt.Printf("\n%s\n\n", buildinfo.String())
	fmt.Printf("I am %s old\n", buildinfo.Age())
}
