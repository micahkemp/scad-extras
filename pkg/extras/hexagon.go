package extras

import "github.com/micahkemp/scad/pkg/scad"

type Hexagon struct {
	Name   string
	Radius float64
}

func (h Hexagon) SCADWriter() scad.SCADWriter {
	return RegularPolygon{
		Name: scad.Name{
			Specified: h.Name,
			Default: "hexagon_component",
		}.String(),
		Sides: 6,
		Radius: h.Radius,
	}.SCADWriter()
}
