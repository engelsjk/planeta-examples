# planeta-examples

Examples for the package [```engelsjk/planeta```](https://github.com/engelsjk/planeta/).

## Basic

```
geomPolygon, _ := geo.ParseGeometry("POLYGON((-2.0 0.0, 0.0 0.0, 0.0 1.0, -2.0 1.0, -2.0 0.0))")
geomLineString, _ := geo.ParseGeometry("LINESTRING(-0.5 0.5, 0.5 0.5)")

print(geomfn.Length(geomLineString))
// 1

print(geomfn.Area(geomPolygon))
// 2

print(geomfn.Intersects(geomPolygon, geomLineString))
// true
```

## GeoJSON

```
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
```

## Geography

```
geogLineString, _ := geo.ParseGeography("LINESTRING(-109.00463303324997 37.08890778791475,-109.09252365824997 36.90629181998808)")
geogPolygon, _ := geo.ParseGeography("POLYGON((-109.1268559336406 37.04617221507986,-108.9620610117656 37.04617221507986,-108.9620610117656 36.9595315239561,-109.1268559336406 36.9595315239561,-109.1268559336406 37.04617221507986))")

print(geogfn.Length(geogLineString, geogfn.UseSphere))
// 21754.45520583837

print(geogfn.Area(geogPolygon, geogfn.UseSphere))
// 1.4098377428322914e+08

print(geogfn.Intersects(geogLineString, geogPolygon))
// true
```

## Encoding
