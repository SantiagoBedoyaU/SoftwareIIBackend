package service

import (
	"context"
	"time"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"
)

type ReportService struct {
	appointmentRepository port.AppointmentRepository
	userRepository 		  port.UserRepository
}

func NewReportService(appointmentRepository port.AppointmentRepository, userRepository port.UserRepository) *ReportService {
	return &ReportService{appointmentRepository: appointmentRepository, userRepository: userRepository}
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
	now := time.Now()
	truncated_now := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if truncated_now.Equal(endDate) || truncated_now.Before(endDate) {
		return nil, domain.ErrNotValidEndDate
	}
	hours := 23*time.Hour + 59*time.Minute
	endDate = endDate.Add(hours)
	report, err := s.appointmentRepository.GenerateAttendanceReport(ctx, startDate, endDate)
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
	now := time.Now()
	truncated_now := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if truncated_now.Equal(endDate) || truncated_now.Before(endDate) {
		return nil, domain.ErrNotValidEndDate
	}
	hours := 23*time.Hour + 59*time.Minute
	endDate = endDate.Add(hours)
	appointments, err := s.appointmentRepository.GenerateWaitingTimeReport(ctx, startDate, endDate)
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
	report.AveragePerDay = days
	report.DayWithMaxWaitingTime = max_day
	report.DayWithMinWaitingTime = min_day

	return &report, nil
}

func (s *ReportService) GenerateUsersDNIReport(ctx context.Context) (*domain.UsersDNIReport, error) {
	role := ctx.Value("userRole").(float64)
	var report domain.UsersDNIReport
	// Only an admin can view reports
	if role != float64(domain.AdminRole) {
		return nil, domain.ErrNotAnAdminRole
	}
	total_cc, total_ti, total_tp, err := s.userRepository.GenerateUsersDNIReport(ctx)
	if err != nil {
		return nil, err
	}
	total_users := total_cc + total_ti + total_tp
	// Counts
	report.TotalUsers = total_users
	report.CCUsers = total_cc
	report.TIUsers = total_ti
	report.TPUsers = total_tp
	// Percentages
	report.CCPercentage = (float64(total_cc * 100) / float64(total_users))
	report.TIPercentage = (float64(total_ti * 100) / float64(total_users))
	report.TPPercentage = (float64(total_tp * 100) / float64(total_users))
	return &report, nil
}

func (s *ReportService) GenerateMostConsultedDoctorsReport(ctx context.Context, startDate, endDate time.Time) (*domain.ConsultedDoctors, error) {
	role := ctx.Value("userRole").(float64)
	var report domain.ConsultedDoctors
	// Only an admin can view reports
	if role != float64(domain.AdminRole) {
		return nil, domain.ErrNotAnAdminRole
	}
	if endDate.Before(startDate) {
		return nil, domain.ErrNotValidDates
	}
	duration := 23*time.Hour + 59*time.Minute
	endDate = endDate.Add(duration)
	now := time.Now()
	truncated_now := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if truncated_now.Equal(endDate) || truncated_now.Before(endDate) {
		return nil, domain.ErrNotValidEndDate
	}
	appointments, err := s.appointmentRepository.GetAppointmentsBetweenDates(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}
	doctors := make(map[string]int)
	// Agrupando y contando
	for _, obj := range appointments {
		doctors[obj.DoctorID]++
	}
	report.Doctors = doctors
	return &report, nil
}
