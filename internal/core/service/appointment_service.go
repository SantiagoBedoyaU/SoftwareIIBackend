package service

import (
	"context"
	"fmt"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"
	"time"
)

type AppointmentService struct {
	appointmentRepository port.AppointmentRepository
	UserService           port.UserService
	emailService          port.EmailService
}

func NewAppointmentService(appointmentRepository port.AppointmentRepository, UserService port.UserService, emailService port.EmailService) *AppointmentService {
	return &AppointmentService{
		appointmentRepository: appointmentRepository,
		UserService:           UserService,
		emailService:          emailService,
	}
}

func (s *AppointmentService) GetByDateRange(ctx context.Context, startDate, endDate time.Time, doctorID, patientID string) ([]domain.Appointment, error) {
	return s.appointmentRepository.GetByDateRange(ctx, startDate, endDate, doctorID, patientID)
}

func (s *AppointmentService) AddAppointmentProcedure(ctx context.Context, appointmentID string, appointmentPatch domain.AppointmentPatch) error {
	return s.appointmentRepository.AddAppointmentProcedure(ctx, appointmentID, appointmentPatch)
}

func (s *AppointmentService) GetHistoryByUser(ctx context.Context, userDNI string) ([]domain.Appointment, error) {
	return s.appointmentRepository.GetHistoryByUser(ctx, userDNI)
}

func (s *AppointmentService) CreateAppointment(ctx context.Context, appointment *domain.Appointment) error {
	// we get the user in order to send an email for the new appointment
	user, err := s.UserService.GetUser(ctx, appointment.PatientID)
	if err != nil {
		return err
	}
	doctor, err := s.UserService.GetUser(ctx, appointment.DoctorID)
	if err != nil {
		return err
	}

	// we can't allow an appointment with a doctor without the appropiate rol
	if doctor.Role != domain.MedicRole {
		return domain.ErrNotAMedicRole
	}
	endDate := appointment.StartDate.Add(15 * time.Minute)
	appointment.EndDate = endDate
	appointment.RealStartDate = appointment.StartDate
	// we can't create two appointments with the same date
	if appointments, _ := s.appointmentRepository.GetByDateRange(ctx, appointment.StartDate, appointment.EndDate, "", appointment.PatientID); len(appointments) > 0 {
		return domain.ErrAlreadyHaveAnAppointment
	}
	if appointment.EndDate.Before(appointment.StartDate) {
		return domain.ErrNotValidDates
	}
	_ = s.emailService.SendAppointmentEmail(ctx, fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email, appointment.StartDate)

	return s.appointmentRepository.CreateAppointment(ctx, appointment)
}

func (s *AppointmentService) CancelAppointment(ctx context.Context, id string) error {
	return s.appointmentRepository.CancelAppointment(ctx, id)
}
