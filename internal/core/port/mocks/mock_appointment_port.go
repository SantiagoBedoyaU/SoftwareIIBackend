// Code generated by MockGen. DO NOT EDIT.
// Source: appointment_port.go
//
// Generated by this command:
//
//	mockgen -source=appointment_port.go -destination=mocks/mock_appointment_port.go -typed
//

// Package mock_port is a generated GoMock package.
package mock_port

import (
	context "context"
	reflect "reflect"
	domain "softwareIIbackend/internal/core/domain"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockAppoitmentService is a mock of AppoitmentService interface.
type MockAppoitmentService struct {
	ctrl     *gomock.Controller
	recorder *MockAppoitmentServiceMockRecorder
	isgomock struct{}
}

// MockAppoitmentServiceMockRecorder is the mock recorder for MockAppoitmentService.
type MockAppoitmentServiceMockRecorder struct {
	mock *MockAppoitmentService
}

// NewMockAppoitmentService creates a new mock instance.
func NewMockAppoitmentService(ctrl *gomock.Controller) *MockAppoitmentService {
	mock := &MockAppoitmentService{ctrl: ctrl}
	mock.recorder = &MockAppoitmentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppoitmentService) EXPECT() *MockAppoitmentServiceMockRecorder {
	return m.recorder
}

// CancelAppointment mocks base method.
func (m *MockAppoitmentService) CancelAppointment(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelAppointment", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelAppointment indicates an expected call of CancelAppointment.
func (mr *MockAppoitmentServiceMockRecorder) CancelAppointment(ctx, id any) *MockAppoitmentServiceCancelAppointmentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelAppointment", reflect.TypeOf((*MockAppoitmentService)(nil).CancelAppointment), ctx, id)
	return &MockAppoitmentServiceCancelAppointmentCall{Call: call}
}

// MockAppoitmentServiceCancelAppointmentCall wrap *gomock.Call
type MockAppoitmentServiceCancelAppointmentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAppoitmentServiceCancelAppointmentCall) Return(arg0 error) *MockAppoitmentServiceCancelAppointmentCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAppoitmentServiceCancelAppointmentCall) Do(f func(context.Context, string) error) *MockAppoitmentServiceCancelAppointmentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAppoitmentServiceCancelAppointmentCall) DoAndReturn(f func(context.Context, string) error) *MockAppoitmentServiceCancelAppointmentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateAppointment mocks base method.
func (m *MockAppoitmentService) CreateAppointment(ctx context.Context, appointment *domain.Appointment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAppointment", ctx, appointment)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAppointment indicates an expected call of CreateAppointment.
func (mr *MockAppoitmentServiceMockRecorder) CreateAppointment(ctx, appointment any) *MockAppoitmentServiceCreateAppointmentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAppointment", reflect.TypeOf((*MockAppoitmentService)(nil).CreateAppointment), ctx, appointment)
	return &MockAppoitmentServiceCreateAppointmentCall{Call: call}
}

// MockAppoitmentServiceCreateAppointmentCall wrap *gomock.Call
type MockAppoitmentServiceCreateAppointmentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAppoitmentServiceCreateAppointmentCall) Return(arg0 error) *MockAppoitmentServiceCreateAppointmentCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAppoitmentServiceCreateAppointmentCall) Do(f func(context.Context, *domain.Appointment) error) *MockAppoitmentServiceCreateAppointmentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAppoitmentServiceCreateAppointmentCall) DoAndReturn(f func(context.Context, *domain.Appointment) error) *MockAppoitmentServiceCreateAppointmentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByDateRange mocks base method.
func (m *MockAppoitmentService) GetByDateRange(ctx context.Context, startDate, endDate time.Time, doctorID, patientID string) ([]domain.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDateRange", ctx, startDate, endDate, doctorID, patientID)
	ret0, _ := ret[0].([]domain.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDateRange indicates an expected call of GetByDateRange.
func (mr *MockAppoitmentServiceMockRecorder) GetByDateRange(ctx, startDate, endDate, doctorID, patientID any) *MockAppoitmentServiceGetByDateRangeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDateRange", reflect.TypeOf((*MockAppoitmentService)(nil).GetByDateRange), ctx, startDate, endDate, doctorID, patientID)
	return &MockAppoitmentServiceGetByDateRangeCall{Call: call}
}

// MockAppoitmentServiceGetByDateRangeCall wrap *gomock.Call
type MockAppoitmentServiceGetByDateRangeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAppoitmentServiceGetByDateRangeCall) Return(arg0 []domain.Appointment, arg1 error) *MockAppoitmentServiceGetByDateRangeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAppoitmentServiceGetByDateRangeCall) Do(f func(context.Context, time.Time, time.Time, string, string) ([]domain.Appointment, error)) *MockAppoitmentServiceGetByDateRangeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAppoitmentServiceGetByDateRangeCall) DoAndReturn(f func(context.Context, time.Time, time.Time, string, string) ([]domain.Appointment, error)) *MockAppoitmentServiceGetByDateRangeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockAppointmentRepository is a mock of AppointmentRepository interface.
type MockAppointmentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAppointmentRepositoryMockRecorder
	isgomock struct{}
}

// MockAppointmentRepositoryMockRecorder is the mock recorder for MockAppointmentRepository.
type MockAppointmentRepositoryMockRecorder struct {
	mock *MockAppointmentRepository
}

// NewMockAppointmentRepository creates a new mock instance.
func NewMockAppointmentRepository(ctrl *gomock.Controller) *MockAppointmentRepository {
	mock := &MockAppointmentRepository{ctrl: ctrl}
	mock.recorder = &MockAppointmentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppointmentRepository) EXPECT() *MockAppointmentRepositoryMockRecorder {
	return m.recorder
}

// CancelAppointment mocks base method.
func (m *MockAppointmentRepository) CancelAppointment(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelAppointment", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelAppointment indicates an expected call of CancelAppointment.
func (mr *MockAppointmentRepositoryMockRecorder) CancelAppointment(ctx, id any) *MockAppointmentRepositoryCancelAppointmentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelAppointment", reflect.TypeOf((*MockAppointmentRepository)(nil).CancelAppointment), ctx, id)
	return &MockAppointmentRepositoryCancelAppointmentCall{Call: call}
}

// MockAppointmentRepositoryCancelAppointmentCall wrap *gomock.Call
type MockAppointmentRepositoryCancelAppointmentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAppointmentRepositoryCancelAppointmentCall) Return(arg0 error) *MockAppointmentRepositoryCancelAppointmentCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAppointmentRepositoryCancelAppointmentCall) Do(f func(context.Context, string) error) *MockAppointmentRepositoryCancelAppointmentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAppointmentRepositoryCancelAppointmentCall) DoAndReturn(f func(context.Context, string) error) *MockAppointmentRepositoryCancelAppointmentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateAppointment mocks base method.
func (m *MockAppointmentRepository) CreateAppointment(ctx context.Context, appointment *domain.Appointment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAppointment", ctx, appointment)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAppointment indicates an expected call of CreateAppointment.
func (mr *MockAppointmentRepositoryMockRecorder) CreateAppointment(ctx, appointment any) *MockAppointmentRepositoryCreateAppointmentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAppointment", reflect.TypeOf((*MockAppointmentRepository)(nil).CreateAppointment), ctx, appointment)
	return &MockAppointmentRepositoryCreateAppointmentCall{Call: call}
}

// MockAppointmentRepositoryCreateAppointmentCall wrap *gomock.Call
type MockAppointmentRepositoryCreateAppointmentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAppointmentRepositoryCreateAppointmentCall) Return(arg0 error) *MockAppointmentRepositoryCreateAppointmentCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAppointmentRepositoryCreateAppointmentCall) Do(f func(context.Context, *domain.Appointment) error) *MockAppointmentRepositoryCreateAppointmentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAppointmentRepositoryCreateAppointmentCall) DoAndReturn(f func(context.Context, *domain.Appointment) error) *MockAppointmentRepositoryCreateAppointmentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByDateRange mocks base method.
func (m *MockAppointmentRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time, doctorID, patientID string) ([]domain.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDateRange", ctx, startDate, endDate, doctorID, patientID)
	ret0, _ := ret[0].([]domain.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDateRange indicates an expected call of GetByDateRange.
func (mr *MockAppointmentRepositoryMockRecorder) GetByDateRange(ctx, startDate, endDate, doctorID, patientID any) *MockAppointmentRepositoryGetByDateRangeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDateRange", reflect.TypeOf((*MockAppointmentRepository)(nil).GetByDateRange), ctx, startDate, endDate, doctorID, patientID)
	return &MockAppointmentRepositoryGetByDateRangeCall{Call: call}
}

// MockAppointmentRepositoryGetByDateRangeCall wrap *gomock.Call
type MockAppointmentRepositoryGetByDateRangeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAppointmentRepositoryGetByDateRangeCall) Return(arg0 []domain.Appointment, arg1 error) *MockAppointmentRepositoryGetByDateRangeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAppointmentRepositoryGetByDateRangeCall) Do(f func(context.Context, time.Time, time.Time, string, string) ([]domain.Appointment, error)) *MockAppointmentRepositoryGetByDateRangeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAppointmentRepositoryGetByDateRangeCall) DoAndReturn(f func(context.Context, time.Time, time.Time, string, string) ([]domain.Appointment, error)) *MockAppointmentRepositoryGetByDateRangeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
