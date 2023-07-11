package entity

type Seat struct {
	ProjectionId int  `json:"projectionId"`
	RowNumber    int  `json:"rowNumber"`
	ColumnNumber int  `json:"columnNumber"`
	Occupied     bool `json:"occupied"`
}
