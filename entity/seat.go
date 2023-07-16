package entity

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	ReservationId int64 `json:"projection"`
	RowNumber     int64 `json:"rowNumber"`
	ColumnNumber  int64 `json:"columnNumber"`
	Occupied      bool  `json:"occupied"`
}
