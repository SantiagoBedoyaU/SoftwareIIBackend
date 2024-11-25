package domain

import "time"

type Procedure struct {
	RealStartDate time.Time `json:"real_start_date" bson:"real_start_date"`
	Description   string 	`json:"description" bson:"description"`
}
