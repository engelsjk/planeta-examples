package main

import (
	"fmt"

	"github.com/engelsjk/planeta/geo"
	"github.com/engelsjk/planeta/geo/geogfn"
)

func main() {

	geogLineString, _ := geo.ParseGeography("LINESTRING(-109.00463303324997 37.08890778791475,-109.09252365824997 36.90629181998808)")
	geogPolygon, _ := geo.ParseGeography("POLYGON((-109.1268559336406 37.04617221507986,-108.9620610117656 37.04617221507986,-108.9620610117656 36.9595315239561,-109.1268559336406 36.9595315239561,-109.1268559336406 37.04617221507986))")

	print(geogfn.Length(geogLineString, geogfn.UseSphere))
	// 21754.45520583837

	print(geogfn.Area(geogPolygon, geogfn.UseSphere))
	// 1.4098377428322914e+08

	print(geogfn.Intersects(geogLineString, geogPolygon))
	// true

}

func print(i interface{}, err error) {
	fmt.Println(i)
}
