package domain

import "time"

type Procedure struct {
	Description   string 	`json:"description" bson:"description"`
}

type AppointmentPatch struct{
	Procedure 	  Procedure	`json:"procedure"`
	RealStartDate time.Time `json:"real_start_date"`
}
