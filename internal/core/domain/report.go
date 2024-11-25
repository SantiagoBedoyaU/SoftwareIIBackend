package domain

import "time"

type AttendanceReport struct {
	TotalPatients           int64 `json:"total_patients" bson:"total_patients"`
	NonAttendingPatients    int64 `json:"non-attending_patients" bson:"non-attending_patients"`
	AttendingPatients       int64 `json:"attending_patients" bson:"attending_patients"`
	AttendancePercentage    int64 `json:"attendance_percentage" bson:"attendance_percentage"`
	NonAttendancePercentage int64 `json:"non-attendance_percentage" bson:"non-attendance_percentage"`
}

type WaitingTimeReport struct {
	AverageWaitingTime    float64            	`json:"average_waiting_time" bson:"average_waiting_time"`
	DaysPerAverage        map[time.Time]float64 `json:"days_per_average" bson:"days_per_average"`
	DayWithMaxWaitingTime time.Time          	`json:"days_with_max_waiting_time" bson:"days_with_max_waiting_time"`
	DayWithMinWaitingTime time.Time          	`json:"days_with_min_waiting_time" bson:"days_with_min_waiting_time"`
}
