package domain

import "time"

type AttendanceReport struct {
	TotalPatients           int64   `json:"total_patients"`
	NonAttendingPatients    int64   `json:"non-attending_patients"`
	AttendingPatients       int64   `json:"attending_patients"`
	AttendancePercentage    float64 `json:"attendance_percentage"`
	NonAttendancePercentage float64 `json:"non-attendance_percentage"`
}

type WaitingTimeReport struct {
	AverageWaitingTime    float64            	`json:"average_waiting_time"`
	AveragePerDay         map[time.Time]float64 `json:"average_per_day"`
	DayWithMaxWaitingTime time.Time          	`json:"days_with_max_waiting_time"`
	DayWithMinWaitingTime time.Time          	`json:"days_with_min_waiting_time"`
}

type UsersDNIReport struct {
	TotalUsers   int64	 `json:"total_users"`
	CCUsers		 int64	 `json:"cc_users"`
	TIUsers		 int64	 `json:"ti_users"`
	TPUsers		 int64	 `json:"tp_users"`
	CCPercentage float64 `json:"cc_percentage"`
	TIPercentage float64 `json:"ti_percentage"`
	TPPercentage float64 `json:"tp_percentage"`
}

type ConsultedDoctors struct {
	Doctors map[string]int `json:"doctors"`
}
