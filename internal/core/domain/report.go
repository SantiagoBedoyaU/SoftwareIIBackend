package domain

import "time"

type AttendanceReport struct {
	TotalPatients           int64 `json:"total_patients" bson:"total_patients"`
	NonAttendingPatients    int64 `json:"non-attending_patients" bson:"non-attending_patients"`
	AttendingPatients       int64 `json:"attending_patients" bson:"attending_patients"`
	AttendancePercentage    float64 `json:"attendance_percentage" bson:"attendance_percentage"`
	NonAttendancePercentage float64 `json:"non-attendance_percentage" bson:"non-attendance_percentage"`
}

type WaitingTimeReport struct {
	AverageWaitingTime    float64            	`json:"average_waiting_time" bson:"average_waiting_time"`
	AveragePerDay         map[time.Time]float64 `json:"average_per_day" bson:"average_per_day"`
	DayWithMaxWaitingTime time.Time          	`json:"days_with_max_waiting_time" bson:"days_with_max_waiting_time"`
	DayWithMinWaitingTime time.Time          	`json:"days_with_min_waiting_time" bson:"days_with_min_waiting_time"`
}
