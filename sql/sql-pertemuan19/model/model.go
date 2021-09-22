package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name  string
	Order []Order
}

type Order struct {
	gorm.Model
	OrderDate  time.Time
	CustomerId uint
	Shipment   []Shipment
}

type Shipment struct {
	gorm.Model
	ShipmentDate time.Time
	OrderId      uint
}
