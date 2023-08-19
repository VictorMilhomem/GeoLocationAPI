package models

type Point struct {
	Type  string    `json:"type"`
	Coord []float64 `json:"coordinates"`
}

type MultiPolygon struct {
	Type   string           `json:"type"`
	Coords [][][][2]float64 `json:"coordinates"`
}

type Store struct {
	Id           int          `json:"id"`
	TradingName  string       `json:"tradingName"`
	OwnerName    string       `json:"ownerName"`
	Document     string       `json:"document"`
	CoverageArea MultiPolygon `json:"coverageArea"`
	Addrs        Point        `json:"addrs"`
}

func NewPoint(coords []float64) *Point {
	return &Point{
		Type:  "Point",
		Coord: coords,
	}
}

func NewMultiPolygon(coords [][][][2]float64) *MultiPolygon {
	return &MultiPolygon{
		Type:   "MultiPolygon",
		Coords: coords,
	}
}

func NewStore(id int, tradingName string, ownerName string, document string, coverageArea MultiPolygon, addrs Point) *Store {
	return &Store{
		Id:           id,
		TradingName:  tradingName,
		OwnerName:    ownerName,
		Document:     document,
		CoverageArea: coverageArea,
		Addrs:        addrs,
	}
}
