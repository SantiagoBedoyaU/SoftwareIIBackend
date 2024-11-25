package service

import (
	"context"
	"time"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"
)

type ReportService struct {
	repo port.AppointmentRepository
}

func NewReportService(repo port.AppointmentRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GenerateAttendanceReport(ctx context.Context, startDate, endDate time.Time) (*domain.AttendanceReport, error) {
	role := ctx.Value("userRole").(float64)
	// Only an admin can view reports
	if role != float64(domain.AdminRole) {
		return nil, domain.ErrNotAnAdminRole
	}
	if endDate.Before(startDate) {
		return nil, domain.ErrNotValidDates
	}
	report, err := s.repo.GenerateAttendanceReport(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return report, nil
}


func (s *ReportService) GenerateWaitingTimeReport(ctx context.Context, startDate, endDate time.Time) (*domain.WaitingTimeReport, error) {
	var report domain.WaitingTimeReport
	var duration time.Duration
	var minutes float64
	var day time.Time
	var max_day, min_day time.Time
	var max_value, min_value float64
	isFirst := true
	
	role := ctx.Value("userRole").(float64)
	// Only an admin can view reports
	if role != float64(domain.AdminRole) {
		return nil, domain.ErrNotAnAdminRole
	}
	if endDate.Before(startDate) {
		return nil, domain.ErrNotValidDates
	}
	appointments, err := s.repo.GenerateWaitingTimeReport(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}
	days := make(map[time.Time]float64)
	if len(appointments) > 0 {
		appointmentsPerDay := make(map[time.Time]float64)
	
		for _, val := range appointments {
			// Average waiting time
			duration = val.RealStartDate.Sub(val.StartDate)
			minutes += float64(duration.Minutes())
			// Group days
			day = time.Date(val.StartDate.Year(), val.StartDate.Month(), val.StartDate.Day(), 0, 0, 0, 0, time.UTC)
			days[day] += float64(duration.Minutes())
			appointmentsPerDay[day] += 1
		}
		prom_minutes := minutes / float64(len(appointments))
		// Search for the max and min average waiting time
		for day, average := range days {
			days[day] /= appointmentsPerDay[day]
			if isFirst {
				max_value, min_value = average, average
				max_day, min_day = day, day
				isFirst = false
			} else {
				if average > max_value {
					max_value = average
					max_day = day
				}
				if average < min_value {
					min_value = average
					min_day = day
				}
			}
		}
		report.AverageWaitingTime = prom_minutes
	} else {
		report.AverageWaitingTime = 0
	}
	report.DaysPerAverage = days
	report.DayWithMaxWaitingTime = max_day
	report.DayWithMinWaitingTime = min_day

	return &report, nil
}
