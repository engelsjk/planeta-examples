package main

import (
	"fmt"

	"github.com/engelsjk/planeta/geo"
	"github.com/engelsjk/planeta/geo/geomfn"
	"github.com/twpayne/go-geom/encoding/geojson"
)

func main() {
	b1 := []byte(`{"type": "Feature","properties": {},"geometry": {"type": "Polygon", "coordinates": [[[-83.5345458984375,39.5633531658293], [-82.4139404296875,39.5633531658293], [-82.4139404296875,40.39258071969131], [-83.5345458984375,40.39258071969131], [-83.5345458984375,39.5633531658293]]]}}`)

	b2 := []byte(`{"type": "Feature","properties": {"name": "example"},"geometry": {"type": "Polygon","coordinates": [[[-83.023681640625,39.104488809440475],[-81.968994140625,39.104488809440475],[-81.968994140625,39.85072092501597],[-83.023681640625,39.85072092501597],[-83.023681640625,39.104488809440475]]]}}`)

	var feature1, feature2 geojson.Feature

	feature1.UnmarshalJSON(b1)
	feature2.UnmarshalJSON(b2)

	geometry1, _ := geo.MakeGeometryFromGeomT(feature1.Geometry)
	geometry2, _ := geo.MakeGeometryFromGeomT(feature2.Geometry)

	intersects, _ := geomfn.Intersects(geometry1, geometry2)

	fmt.Println(intersects)
	// true

	intersection, _ := geomfn.Intersection(geometry1, geometry2)

	g, _ := intersection.AsGeomT()
	b, _ := geojson.Marshal(g)

	fmt.Println(string(b))
	// {"type":"Polygon","coordinates":[[[-82.4139404296875,39.85072092501597],[-82.4139404296875,39.5633531658293],[-83.023681640625,39.5633531658293],[-83.023681640625,39.85072092501597],[-82.4139404296875,39.85072092501597]]]}
}
