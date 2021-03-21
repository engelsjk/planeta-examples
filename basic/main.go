package main

import (
	"fmt"

	"github.com/engelsjk/planeta/geo"
	"github.com/engelsjk/planeta/geo/geomfn"
)

func main() {

	geomPolygon, _ := geo.ParseGeometry("POLYGON((-2.0 0.0, 0.0 0.0, 0.0 1.0, -2.0 1.0, -2.0 0.0))")
	geomLineString, _ := geo.ParseGeometry("LINESTRING(-0.5 0.5, 0.5 0.5)")

	print(geomfn.Length(geomLineString))
	// 1

	print(geomfn.Area(geomPolygon))
	// 2

	print(geomfn.Intersects(geomPolygon, geomLineString))
	// true
}

func print(i interface{}, err error) {
	fmt.Println(i)
}
