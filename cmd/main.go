package main

import (
	"fmt"

	ext "github.com/djdhm/external-dep/lib"
	dep1 "github.com/djdhm/moab-dep-1/lib"
	dep2 "github.com/djdhm/moab-dep-2/lib"
)

func main() {
	dep1.PrintVersion()
	dep2.TestLib()
	fmt.Println("Direct call from project to external dep with version = " + ext.Version())

}
