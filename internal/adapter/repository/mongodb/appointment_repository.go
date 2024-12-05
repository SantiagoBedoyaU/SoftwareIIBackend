package mongodb

import (
	"context"
	"errors"
	"softwareIIbackend/internal/core/domain"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppointmentRepository struct {
	CollName string
	conn     *MongoDBConnection
}

func NewAppointmentRepository(collname string, conn *MongoDBConnection) *AppointmentRepository {
	return &AppointmentRepository{CollName: collname, conn: conn}
}

func (r *AppointmentRepository) AddAppointmentProcedure(ctx context.Context, appointmentID string, appointmentPatch domain.AppointmentPatch) error {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	objID, err := primitive.ObjectIDFromHex(appointmentID)
	if err != nil {
		return domain.ErrInvalidIDFormat
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"real_start_date": appointmentPatch.RealStartDate,
			"status": domain.AppointmentStatusDone,
		},
		"$push": bson.M{
			"procedures": appointmentPatch.Procedure,
		},
	}
	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrAppointmentNotFound
		}
		return err
	}
	return nil
}

func (r *AppointmentRepository) GetHistoryByUser(ctx context.Context, userDNI string) ([]domain.Appointment, error) {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	now := time.Now()
	filter := bson.M{
		"patient_id": userDNI,
		"end_date": bson.M{
			"$lt": now,
		},
	}
	results, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	appointments := make([]domain.Appointment, 0)
	if err := results.All(ctx, &appointments); err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time, doctorID, patientID string) ([]domain.Appointment, error) {
	coll := r.conn.GetDatabase().Collection(r.CollName)

	filter := bson.M{
		"$and": []bson.M{
			{
				"$or": []bson.M{
					{
						"start_date": bson.M{"$lt": endDate},
						"end_date":   bson.M{"$gt": startDate},
					},
					{
						"start_date": bson.M{"$lt": endDate},
						"end_date":   bson.M{"$gt": startDate},
					},
					{
						"start_date": bson.M{"$gte": startDate},
						"end_date":   bson.M{"$lte": endDate},
					},
				},
			},
			{
				"status": bson.M{"$ne": domain.AppointmentStatusCancelled},
			},
		},
	}
	doctorID = strings.TrimSpace(doctorID)
	if doctorID != "" {
		filter["doctor_id"] = doctorID
	}

	patientID = strings.TrimSpace(patientID)
	if patientID != "" {
		filter["patient_id"] = patientID
	}

	results, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	appointments := make([]domain.Appointment, 0)
	if err := results.All(ctx, &appointments); err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) CreateAppointment(ctx context.Context, appointment *domain.Appointment) error {
	coll := r.conn.GetDatabase().Collection(r.CollName)

	result, err := coll.InsertOne(ctx, appointment)
	appointment.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return err
}

func (r *AppointmentRepository) CancelAppointment(ctx context.Context, id string) error {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrInvalidIDFormat
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{
		"$set": bson.M{
			"status": domain.AppointmentStatusCancelled,
		},
	}
	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrAppointmentNotFound
		}
		return err
	}
	return nil
}

func (r *AppointmentRepository) GenerateAttendanceReport(ctx context.Context, startDate, endDate time.Time) (*domain.AttendanceReport, error){
	coll := r.conn.GetDatabase().Collection(r.CollName)

	var report domain.AttendanceReport
	filter := bson.M{
		"$or": []bson.M{
			{
				"start_date": bson.M{"$gte": startDate},
				"end_date":   bson.M{"$lte": endDate},
			},
		},
			
	}
	total_patients, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	filter["status"] = domain.AppointmentStatusPending

	non_attending_patients, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	attending_patients := total_patients - non_attending_patients

	report.TotalPatients = total_patients
	report.AttendingPatients = attending_patients
	report.NonAttendingPatients = non_attending_patients
	if total_patients > 0 {
		report.AttendancePercentage = (float64(attending_patients) * 100) / float64(total_patients)
		report.NonAttendancePercentage = (float64(non_attending_patients) * 100) / float64(total_patients)
	} else {
		report.AttendancePercentage = 0
		report.NonAttendancePercentage = 0
	}
	return &report, nil
}

func (r *AppointmentRepository) GenerateWaitingTimeReport(ctx context.Context, startDate, endDate time.Time) ([]domain.Appointment, error) {
	coll := r.conn.GetDatabase().Collection(r.CollName)

	filter := bson.M{
		"$and": []bson.M{
			{
				"$or": []bson.M{
					{
						"start_date": bson.M{"$gte": startDate},
						"end_date":   bson.M{"$lte": endDate},
					},
				},
			},
			{
				"status": bson.M{"$eq": domain.AppointmentStatusDone},
			},
		},
	}
	results, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	appointments := make([]domain.Appointment, 0)

	if err := results.All(ctx, &appointments); err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) GetAppointmentsBetweenDates(ctx context.Context, startDate, endDate time.Time) ([]domain.Appointment, error) {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	
	filter := bson.M{
		"$and": []bson.M{
			{
				"$or": []bson.M{
					{
						"start_date": bson.M{"$lt": endDate},
						"end_date":   bson.M{"$gt": startDate},
					},
					{
						"start_date": bson.M{"$gte": startDate},
						"end_date":   bson.M{"$lte": endDate},
					},
				},
			},
			{
				"status": bson.M{"$eq": domain.AppointmentStatusDone},
			},
		},
	}
	results, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	appointments := make([]domain.Appointment, 0)
	if err := results.All(ctx, &appointments); err != nil {
		return nil, err
	}
	return appointments, nil
}
