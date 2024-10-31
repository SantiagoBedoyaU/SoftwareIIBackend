package mongodb

import (
	"context"
	"softwareIIbackend/internal/core/domain"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppointmentRepository struct {
	CollName string
	conn     *MongoDBConnection
}

func NewAppointmentRepository(collname string, conn *MongoDBConnection) *AppointmentRepository {
	return &AppointmentRepository{CollName: collname, conn: conn}
}

func (r *AppointmentRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time, doctorID, patientID string) ([]domain.Appointment, error) {
	coll := r.conn.GetDatabase().Collection(r.CollName)

	filter := bson.M{
		"start_date": bson.M{
			"$gte": startDate,
		},
		"end_date": bson.M{
			"$lte": endDate,
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