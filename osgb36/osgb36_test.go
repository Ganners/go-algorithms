package osgb36

import (
	"math"
	"testing"
)

func TestBSGToLatLong(t *testing.T) {
	// Test cases, base results have been grabbed using this converter:
	// http://www.bgs.ac.uk/data/webservices/convertForm.cfm
	testCases := []struct {
		Easting  float64
		Northing float64
		Lat      float64
		Long     float64
	}{
		{
			535267, 181084,
			51.512497043, -0.0520978712712,
		},
		{
			534553, 181391,
			51.5154261692, -0.0622637053248,
		},
		{
			522310, 176034,
			51.4700623957, -0.240465983756,
		},
	}

	for _, tst := range testCases {
		lat, lon := BSGToLatLong(tst.Northing, tst.Easting)

		// Allowed error..
		acceptableDistance := 0.00003

		// Test if they match
		dLat := math.Abs(lat - tst.Lat)
		dLon := math.Abs(lon - tst.Long)

		if dLat > acceptableDistance || dLon > acceptableDistance {

			t.Errorf("Latitude and longitude is too far from expected: Lat:%f Lon:%f",
				dLat, dLon)
		}
	}
}
