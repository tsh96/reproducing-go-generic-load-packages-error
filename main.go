package main

import (
	_ "github.com/tsh96/reproducing-go-generic-load-packages-error/g"

	"golang.org/x/tools/go/packages"
)

func main() {
	mode := packages.NeedTypes | packages.NeedTypesInfo
	packages.Load(&packages.Config{Mode: mode}, ".")
}
