package mongodb

import (
	"context"
	"softwareIIbackend/internal/core/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type AppointmentRepository struct {
	CollName string
	conn     *MongoDBConnection
}

func NewAppointmentRepository(collname string, conn *MongoDBConnection) *AppointmentRepository {
	return &AppointmentRepository{CollName: collname, conn: conn}
}

func (r *AppointmentRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]domain.Appointment, error) {
	coll := r.conn.GetDatabase().Collection(r.CollName)

	filter := bson.M{
		"start_date": bson.M{
			"$gte": startDate,
		},
		"end_date": bson.M{
			"$lte": endDate,
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
