package entity

type Seat struct {
	ID            int64 `json:"id"`
	ReservationId int64 `json:"projection"`
	RowNumber     int64 `json:"rowNumber"`
	ColumnNumber  int64 `json:"columnNumber"`
	Occupied      bool  `json:"occupied"`
}
