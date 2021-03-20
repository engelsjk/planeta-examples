package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/engelsjk/planeta/geo/geopb"
	"github.com/engelsjk/planeta/geo/geos"
	"github.com/twpayne/go-geom/encoding/ewkb"
	"github.com/twpayne/go-geom/encoding/geojson"
)

func main() {

	s1 := `{
		"type": "Feature",
		"properties": {},
		"geometry": {
		  "type": "Polygon",
		  "coordinates": [
			[
			  [
				-83.5345458984375,
				39.5633531658293
			  ],
			  [
				-82.4139404296875,
				39.5633531658293
			  ],
			  [
				-82.4139404296875,
				40.39258071969131
			  ],
			  [
				-83.5345458984375,
				40.39258071969131
			  ],
			  [
				-83.5345458984375,
				39.5633531658293
			  ]
			]
		  ]
		}
	  }`

	s2 := `{
		"type": "Feature",
		"properties": {
			"name": "example"
		},
		"geometry": {
		  "type": "Polygon",
		  "coordinates": [
			[
			  [
				-83.023681640625,
				39.104488809440475
			  ],
			  [
				-81.968994140625,
				39.104488809440475
			  ],
			  [
				-81.968994140625,
				39.85072092501597
			  ],
			  [
				-83.023681640625,
				39.85072092501597
			  ],
			  [
				-83.023681640625,
				39.104488809440475
			  ]
			]
		  ]
		}
	  }`

	g1, err := featureToEWKB(s1)
	if err != nil {
		log.Fatal(err)
	}

	g2, err := featureToEWKB(s2)
	if err != nil {
		log.Fatal(err)
	}

	// valid

	if isValid, err := geos.IsValid(g1); err == nil {
		fmt.Printf("g1 valid: %t\n", isValid)
	} else {
		log.Fatal(err)
	}

	if isValid, err := geos.IsValid(g2); err == nil {
		fmt.Printf("g2 valid: %t\n", isValid)
	} else {
		log.Fatal(err)
	}

	// intersects

	if intersects, err := geos.Intersects(g1, g2); err == nil {
		fmt.Printf("g1 intersects g2: %t\n", intersects)
	} else {
		log.Fatal(err)
	}

	// intersection

	intersection, err := geos.Intersection(g1, g2)

	g, err := ewkb.Unmarshal(intersection)
	if err != nil {
		log.Fatal(err)
	}
	b, err := geojson.Marshal(g)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("intersection: %s\n", string(b))

}

func featureToEWKB(fs string) (geopb.EWKB, error) {
	f := &geojson.Feature{}
	err := json.Unmarshal([]byte(fs), f)
	if err != nil {
		return nil, err
	}
	b, err := ewkb.Marshal(f.Geometry, binary.BigEndian)
	if err != nil {
		return nil, err
	}
	return geopb.EWKB(b), nil
}
