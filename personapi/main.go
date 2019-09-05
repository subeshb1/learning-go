package main

import (
	"personapi/bindings"

	gy "github.com/graniticio/granitic-yaml/v2"
)

func main() {

	gy.StartGraniticWithYaml(bindings.Components())
}
