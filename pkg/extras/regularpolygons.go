package extras

import (
	"github.com/micahkemp/scad/pkg/scad"
	"log"
	"math"
)

type RegularPolygon struct {
	Name         string
	Sides        int
	Circumradius float64
	Apothem      float64
}

func (r RegularPolygon) angleRadiansPerPoint() float64 {
	return (math.Pi * 2) / float64(r.Sides)
}

func (r RegularPolygon) angleForPointNumber(pointNumber int, rotateRadians float64) float64 {
	return (r.angleRadiansPerPoint() * float64(pointNumber)) - rotateRadians
}

func (r RegularPolygon) pointForPointNumber(pointNumber int, rotateRadians float64) scad.XYCoordinate {
	return scad.XYCoordinate{
		X: r.Circumradius * math.Cos(r.angleForPointNumber(pointNumber, rotateRadians)),
		Y: r.Circumradius * math.Sin(r.angleForPointNumber(pointNumber, rotateRadians)),
	}
}

func (r RegularPolygon) CircumradiusFromApothem() float64 {
	return r.Apothem / (math.Cos(math.Pi / float64(r.Sides)))
}

func (r RegularPolygon) SCADWriter() scad.SCADWriter {
	name := scad.Name{
		Specified: r.Name,
		Default:   "regular_polygon_component",
	}.String()

	if r.Circumradius != 0 && r.Apothem != 0 {
		log.Fatalf("RegularPolygon (%s) defined with both Circumradius (%f) and Apothem (%f)",
			name,
			r.Circumradius,
			r.Apothem,
		)
	}

	pointsPolygon := r
	rotateRadians := 0.0

	if r.Apothem != 0 {
		pointsPolygon.Circumradius = r.CircumradiusFromApothem()
		pointsPolygon.Apothem = 0
		rotateRadians = r.angleRadiansPerPoint() / 2
	}

	points := make(scad.XYCoordinates, r.Sides)

	for pointNumber, _ := range points {
		points[pointNumber] = pointsPolygon.pointForPointNumber(pointNumber, rotateRadians)
	}

	return scad.Polygon{
		Name:   name,
		Points: points,
	}.SCADWriter()
}
