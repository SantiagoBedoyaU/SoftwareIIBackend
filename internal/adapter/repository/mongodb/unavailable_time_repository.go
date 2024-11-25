package mongodb

import (
	"context"
	"errors"
	"softwareIIbackend/internal/core/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UnavailableTimeRepository struct {
	CollName string
	conn     *MongoDBConnection
}

func NewUnavailableTimeRepository(collname string, conn *MongoDBConnection) *UnavailableTimeRepository {
	return &UnavailableTimeRepository{
		CollName: collname,
		conn:     conn,
	}
}

func (r *UnavailableTimeRepository) GetUnavailableTime(ctx context.Context, startDate, endDate time.Time, doctorID string) ([]domain.UnavailableTime, error) {
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
				"doctor_id": bson.M{"$eq": doctorID},
			},
		},
	}

	results, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	uts := make([]domain.UnavailableTime, 0)
	if err := results.All(ctx, &uts); err != nil {
		return nil, err
	}
	return uts, nil
}
func (r *UnavailableTimeRepository) CreateUnavailableTime(ctx context.Context, at *domain.UnavailableTime) error {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	result, err := coll.InsertOne(ctx, at)
	at.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return err
}
func (r *UnavailableTimeRepository) UpdateUnavailableTime(ctx context.Context, at *domain.UnavailableTime) error {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	objID, err := primitive.ObjectIDFromHex(at.ID)
	if err != nil {
		return domain.ErrInvalidIDFormat
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"start_date": at.StartDate,
			"end_date":   at.EndDate,
		},
	}

	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrUnavailableTimeNotFound
		}
		return err
	}
	return nil
}
func (r *UnavailableTimeRepository) DeleteUnavailableTime(ctx context.Context, id string) error {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrInvalidIDFormat
	}
	filter := bson.M{"_id": objID}
	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrUnavailableTimeNotFound
		}
		return err
	}
	return nil
}
