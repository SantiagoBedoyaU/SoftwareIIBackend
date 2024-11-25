package domain

import "time"

type AppointmentStatus int

const (
	AppointmentStatusPending AppointmentStatus = iota
	AppointmentStatusCancelled
	AppointmentStatusDone
)

type Appointment struct {
	ID         	  string            `json:"id" bson:"_id,omitempty"`
	StartDate  	  time.Time         `json:"start_date" bson:"start_date"`
	RealStartDate time.Time         `json:"real_start_date" bson:"real_start_date" default:"2006-01-02T00:00:00Z"`
	EndDate    	  time.Time         `json:"end_date" bson:"end_date"`
	DoctorID   	  string            `json:"doctor_id" bson:"doctor_id"`
	PatientID     string            `json:"patient_id" bson:"patient_id"`
	Status        AppointmentStatus `json:"status" bson:"status"`
	Procedures    []Procedure       `json:"procedures" bson:"procedures,omitempty"`
}
