package examples

import "github.com/micahkemp/scad/pkg/scad"

var Models = scad.Models{
	"decagon":         &exampleDecagon,
	"hexagon":         &exampleHexagon,
	"hexagon_apothem": &exampleHexagonApothem,
}
