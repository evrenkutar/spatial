package spatial

import (
	"encoding/json"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestPointEncodeExact(t *testing.T) {
	precision := 6
	path := []Point{
		Point{
			Lat: 38.796006,
			Lng: -121.314648,
		},
		Point{
			Lat: 38.804000,
			Lng: -121.353282,
		},
	}

	polyline := Encode(path, precision)
	assert.NotEqual(t, "", polyline)
	assert.Equal(t, "km|~hAntmkfFsrNrmjA", polyline)

	points, err := Decode(polyline, precision)
	assert.NoError(t, err)
	assert.Equal(t, path, points)
}

func TestPointEncodingPrecisionRound(t *testing.T) {
	precision := 5
	path := []Point{
		Point{
			Lat: 38.796006,
			Lng: -121.314648,
		},
		Point{
			Lat: 38.804000,
			Lng: -121.353282,
		},
	}
	roundedPath := []Point{
		Point{
			Lat: 38.79601,
			Lng: -121.31465,
		},
		Point{
			Lat: 38.80400,
			Lng: -121.35328,
		},
	}

	polyline := Encode(path, precision)
	assert.NotEqual(t, "", polyline)
	assert.Equal(t, "ajxkFpgmcV}p@lpF", polyline)

	points, err := Decode(polyline, precision)
	assert.NoError(t, err)
	assert.Equal(t, roundedPath, points)
}

func TestPointEncodingLength(t *testing.T) {
	precision := 6
	path := []Point{
		Point{
			Lat: 38.796006,
			Lng: -121.314648,
		},
		Point{
			Lat: 38.804000,
			Lng: -121.353282,
		},
	}

	polyline := Encode(path, precision)
	jsonPoints, err := json.Marshal(path)
	if err != nil {
		assert.Error(t, err)
	}
	assert.True(t, len(polyline) < len(jsonPoints))
}
