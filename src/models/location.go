package models

import (
	"encoding/json"
	"math"

	gormjsonb "github.com/dariubs/gorm-jsonb"
)

type Point struct {
	Type  string    `json:"type"`
	Coord []float64 `json:"coordinates"`
}

type MultiPolygon struct {
	Type   string           `json:"type"`
	Coords [][][][2]float64 `json:"coordinates"`
}

func ConvertJSONBToRawMessage(jsonb gormjsonb.JSONB) (json.RawMessage, error) {
	mjson, err := json.Marshal(jsonb)
	return json.RawMessage(mjson), err
}

func ConvertJSONToPoints(jsonData gormjsonb.JSONB) (*Point, error) {
	var points *Point
	raw, err := ConvertJSONBToRawMessage(jsonData)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(raw, &points); err != nil {
		return nil, err
	}
	return points, nil
}

func (p *Point) EuclideanDistance(p1, p2 Point) float64 {
	if len(p1.Coord) != len(p2.Coord) {
		panic("Points have different dimensions")
	}

	sumOfSquares := 0.0
	for i := 0; i < len(p1.Coord); i++ {
		diff := p1.Coord[i] - p2.Coord[i]
		sumOfSquares += diff * diff
	}

	return math.Sqrt(sumOfSquares)
}
