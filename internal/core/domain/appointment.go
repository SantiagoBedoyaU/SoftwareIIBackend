package domain

import "time"

type AppointmentStatus int

const (
	AppointmentStatusPending AppointmentStatus = iota
	AppointmentStatusCancelled
	AppointmentStatusDone
)

type Appointment struct {
	ID         string            `json:"id" bson:"_id,omitempty"`
	StartDate  time.Time         `json:"start_date" bson:"start_date"`
	EndDate    time.Time         `json:"end_date" bson:"end_date"`
	DoctorID   string            `json:"doctor_id" bson:"doctor_id"`
	PatientID  string            `json:"patient_id" bson:"patient_id"`
	Status     AppointmentStatus `json:"status" bson:"status"`
	Procedures []Procedure       `json:"procedures" bson:"procedures"`
}
