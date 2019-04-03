package osgb36

import "math"

const (
	radToDeg = 180 / math.Pi
	degToRad = math.Pi / 180

	b = 6356256.909
	a = 6377563.396

	f0               = 0.9996012717
	natGridOriginLat = 49 * degToRad
	natGridOriginLon = -2 * degToRad

	northingsOrigin = -100000.0
	eastingsOrigin  = 400000.0

	eccentricity = 1 - (b*b)/(a*a)
	n            = (a - b) / (a + b)
	n2           = n * n
	n3           = n * n * n
)

// Converts BSG (british national grid) or OSGB36 to LAT/LNG (or WGS84)
func BSGToLatLong(northing, easting float64) (float64, float64) {

	lat := natGridOriginLat
	m := 0.0

	for (northing - northingsOrigin - m) >= 1e-5 {

		lat = (northing-northingsOrigin-m)/(a*f0) + lat
		ma := (1 + n + (5/4)*n2 + (5/4)*n3) * (lat - natGridOriginLat)
		mb := (3*n + 3*n*n + (21/8)*n3) * math.Sin(lat-natGridOriginLat) *
			math.Cos(lat+natGridOriginLat)
		mc := ((15/8)*n2 + (15/8)*n3) * math.Sin(2*(lat-natGridOriginLat)) *
			math.Cos(2*(lat+natGridOriginLat))
		md := (35 / 24) * n3 * math.Sin(3*(lat-natGridOriginLat)) * math.Cos(3*(lat+natGridOriginLat))
		m = b * f0 * (ma - mb + mc - md)
	}

	cosLat := math.Cos(lat)
	sinLat := math.Sin(lat)

	nu := a * f0 / math.Sqrt(1-eccentricity*sinLat*sinLat)
	rho := a * f0 * (1 - eccentricity) /
		math.Pow(1-eccentricity*sinLat*sinLat, 1.5)

	eta2 := nu/rho - 1
	tanLat := math.Tan(lat)
	tan2lat := tanLat * tanLat
	tan4lat := tan2lat * tan2lat
	tan6lat := tan4lat * tan2lat
	secLat := 1 / cosLat

	nu3 := nu * nu * nu
	nu5 := nu3 * nu * nu
	nu7 := nu5 * nu * nu

	vii := tanLat / (2 * rho * nu)
	viii := tanLat / (24 * rho * nu3) * (5 + 3*tan2lat + eta2 - 9*tan2lat*eta2)
	ix := tanLat / (720 * rho * nu5) * (61 + 90*tan2lat + 45*tan4lat)
	x := secLat / nu
	xi := secLat / (6 * nu3) * (nu/rho + 2*tan2lat)
	xii := secLat / (120 * nu5) * (5 + 28*tan2lat + 24*tan4lat)
	xiia := secLat / (5040 * nu7) * (61 + 662*tan2lat + 1320*tan4lat + 720*tan6lat)

	de := easting - eastingsOrigin
	de2 := de * de
	de3 := de2 * de
	de4 := de2 * de2
	de5 := de3 * de2
	de6 := de4 * de2
	de7 := de5 * de2

	lat_1 := lat - vii*de2 + viii*de4 - ix*de6
	lon_1 := natGridOriginLon + x*de - xi*de3 + xii*de5 - xiia*de7

	H := 0.0
	x_1 := (nu/f0 + H) * math.Cos(lat_1) * math.Cos(lon_1)
	y_1 := (nu/f0 + H) * math.Cos(lat_1) * math.Sin(lon_1)
	z_1 := ((1-eccentricity)*nu/f0 + H) * math.Sin(lat_1)

	s := -20.4894 * 10e-6
	tx, ty, tz := 446.448, -125.157, 542.060
	rxs, rys, rzs := 0.1502, 0.2470, 0.8421
	rx, ry, rz := rxs*math.Pi/(180*3600.0), rys*math.Pi/(180*3600.0), rzs*math.Pi/(180*3600)

	x_2 := tx + (1+s)*x_1 + (-rz)*y_1 + (ry)*z_1
	y_2 := ty + (rz)*x_1 + (1+s)*y_1 + (-rx)*z_1
	z_2 := tz + (-ry)*x_1 + (rx)*y_1 + (1+s)*z_1

	a_2, b_2 := 6378137.0, 6356752.3141
	e2_2 := 1 - (b_2*b_2)/(a_2*a_2)
	p := math.Sqrt(x_2*x_2 + y_2*y_2)

	lat = math.Atan2(z_2, (p * (1 - e2_2)))

	lon := math.Atan2(y_2, x_2)
	latold := 2 * math.Pi

	for math.Abs(lat-latold) > 10e-16 {
		lat, latold = latold, lat
		latOldSin := math.Sin(latold)
		nu2 := a_2 / math.Sqrt(1-e2_2*(latOldSin*latOldSin))
		lat = math.Atan2(z_2+e2_2*nu2*latOldSin, p)
	}

	lat = lat * radToDeg
	lon = lon * radToDeg

	return lat, lon
}
