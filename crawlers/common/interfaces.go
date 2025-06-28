package common

import (
	"time"
)

type Parser interface {
	Name() string
	Parse(keyword string) ([]Result, error)
}

type Variant struct {
	Price      float64
	Stock      int
	Properties map[string]string
}

type Result struct {
	Title         string
	Discription   string
	Price         float32
	Stock         int
	SoldBy        string
	DeliveryTime  time.Time
	ShipmentPrice float32
	Properties    map[string]string
	Variations    []Variant
	ListingDate   time.Time
	CreatedAt     time.Time
	UpdateAt      time.Time
	Source        string
}

