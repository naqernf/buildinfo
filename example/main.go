package main

import (
	"fmt"

	"github.com/naqernf/buildinfo"
)

func main() {
	fmt.Printf("Hallo World, I am %s old\n", buildinfo.Age())
	fmt.Printf("This build in called example-%s\n", buildinfo.Name())
	fmt.Printf("Is build from commit: %s\n", buildinfo.Githash())
	fmt.Printf("Build time was :%s", buildinfo.Buildstamp())
}
