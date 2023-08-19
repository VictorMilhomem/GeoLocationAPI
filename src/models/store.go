package models

import (
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

type Store struct {
	Id           int             `json:"id"`
	Tradingname  string          `json:"tradingname"`
	Ownername    string          `json:"ownername"`
	Document     string          `json:"document"`
	Coveragearea gormjsonb.JSONB `json:"coveragearea"`
	Addrs        gormjsonb.JSONB `json:"addrs"`
}

func NewStore(id int, tradingName string, ownerName string, document string, coverageArea gormjsonb.JSONB, addrs gormjsonb.JSONB) *Store {
	return &Store{
		Id:           id,
		Tradingname:  tradingName,
		Ownername:    ownerName,
		Document:     document,
		Coveragearea: coverageArea,
		Addrs:        addrs,
	}
}
