package extras

import (
	"github.com/micahkemp/scad/pkg/scad"
	"math"
)

type RegularPolygon struct {
	Name string
	Sides int
	Radius float64
}

func (r RegularPolygon) angleRadiansPerPoint() float64 {
	return (math.Pi * 2) / float64(r.Sides)
}

func (r RegularPolygon) angleForPointNumber(pointNumber int) float64 {
	return r.angleRadiansPerPoint() * float64(pointNumber)
}


func (r RegularPolygon) pointForPointNumber(pointNumber int) scad.XYCoordinate {
	return scad.XYCoordinate{
		X: r.Radius * math.Cos(r.angleForPointNumber(pointNumber)),
		Y: r.Radius * math.Sin(r.angleForPointNumber(pointNumber)),
	}
}

func (r RegularPolygon) SCADWriter() scad.SCADWriter {
	points := make(scad.XYCoordinates, r.Sides)

	for pointNumber, _ := range points {
		points[pointNumber] = r.pointForPointNumber(pointNumber)
	}

	return scad.Polygon{
		Name: scad.Name{
			Specified: r.Name,
			Default: "regular_polygon_component",
		}.String(),
		Points: points,
	}.SCADWriter()
}
