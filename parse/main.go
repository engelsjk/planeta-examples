package main

import (
	"fmt"
	"log"

	"github.com/engelsjk/planeta/geo"
	"github.com/engelsjk/planeta/geo/geos"
)

func main() {

	s1 := `{
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
	  }`

	s2 := `{
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
	  }`

	g1, err := geo.ParseGeometry(s1)
	if err != nil {
		log.Fatal(err)
	}

	g2, err := geo.ParseGeometry(s2)
	if err != nil {
		log.Fatal(err)
	}

	if intersects, err := geos.Intersects(g1.EWKB(), g2.EWKB()); err == nil {
		fmt.Printf("g1 intersects g2: %t\n", intersects)
	} else {
		log.Fatal(err)
	}
}
