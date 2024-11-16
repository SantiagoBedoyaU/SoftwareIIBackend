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
