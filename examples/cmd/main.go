package main

import (
	"github.com/micahkemp/scad-extras/examples/examples"
	"github.com/micahkemp/scad/pkg/scad"
)

func main() {
	scad.CLIHandler(examples.Models)
}
