package mongodb

import (
	"context"
	"errors"
	"softwareIIbackend/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	CollName string
	conn     *MongoDBConnection
}

func NewUserRepository(collname string, conn *MongoDBConnection) *UserRepository {
	return &UserRepository{conn: conn, CollName: collname}
}

func (r *UserRepository) GetUser(ctx context.Context, DNI string) (*domain.User, error) {
	dbname := r.conn.DBName
	coll := r.conn.Client.Database(dbname).Collection(r.CollName)

	var user domain.User
	filter := bson.D{{Key: "dni", Value: DNI}}
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.UserNotFoundErr
		}
		return nil, err
	}
	return &user, nil
}
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	dbname := r.conn.DBName
	coll := r.conn.Client.Database(dbname).Collection(r.CollName)

	var user domain.User
	filter := bson.D{{Key: "email", Value: email}}
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.UserNotFoundErr
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	dbname := r.conn.DBName
	coll := r.conn.Client.Database(dbname).Collection(r.CollName)

	result, err := coll.InsertOne(ctx, user)
	user.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return err
}
