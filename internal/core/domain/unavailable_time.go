package domain

import "time"

type UnavailableTime struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	StartDate time.Time `json:"start_date" bson:"start_date"`
	EndDate   time.Time `json:"end_date" bson:"end_date"`
	DoctorID  string    `json:"doctor_id" bson:"doctor_id"`
}
