package domain

import "time"

type AppointmentStatus int

const (
	AppointmentStatusPending AppointmentStatus = iota
	AppointmentStatusCancelled
	AppointmentStatusDone
)

type Appointment struct {
	ID        string            `json:"id" bson:"_id,omitempty"`
	DateTime  time.Time         `json:"date_time" bson:"date_time"`
	DoctorID  string            `json:"doctor_id" bson:"doctor_id"`
	PatientID string            `json:"patient_id" bson:"patient_id"`
	Status    AppointmentStatus `json:"status" bson:"status"`
}
