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

func (r *UserRepository) GetUser(ctx context.Context, dni string) (*domain.User, error) {
	coll := r.conn.GetDatabase().Collection(r.CollName)

	var user domain.User
	filter := bson.D{{Key: "dni", Value: dni}}
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	coll := r.conn.GetDatabase().Collection(r.CollName)

	var user domain.User
	filter := bson.D{{Key: "email", Value: email}}
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	coll := r.conn.GetDatabase().Collection(r.CollName)

	result, err := coll.InsertOne(ctx, user)
	user.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return err
}

func (r *UserRepository) UpdateUserPassword(ctx context.Context, user *domain.User) error {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	filter := bson.D{{Key: "dni", Value: user.DNI}}
	update := bson.M{
		"$set": bson.M{
			"password": user.Password,
		},
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrUserNotFound
		}
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUserInformation(ctx context.Context, user *domain.User) error {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	filter := bson.D{{Key: "dni", Value: user.DNI}}
	update := bson.M{
		"$set": bson.M{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
			"phone":      user.Phone,
			"address":    user.Address,
		},
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrUserNotFound
		}
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUserRole(ctx context.Context, updateRole *domain.UpdateRole) error {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	filter := bson.D{{Key: "dni", Value: updateRole.DNI}}
	update := bson.M{
		"$set": bson.M{
			"role": updateRole.NewRole,
		},
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrUserNotFound
		}
		return err
	}
	return nil
}

func (r *UserRepository) GetUsersByRole(ctx context.Context, role domain.UserRole) ([]domain.User, error) {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	filter := bson.D{{Key: "role", Value: role}}
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	users := make([]domain.User, 0)
	for cur.Next(ctx) {
		var user domain.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GenerateUsersDNIReport(ctx context.Context) (int64, int64, int64, error) {
	coll := r.conn.GetDatabase().Collection(r.CollName)
	// Count the CC users
	filter := bson.M{
		"type_dni": domain.TypeDniCC,
	}
	total_CC_users, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return 0, 0, 0, err
	}
	// Count the TI users
	filter = bson.M{
		"type_dni": domain.TypeDniTI,
	}
	total_TI_users, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return 0, 0, 0, err
	}
	// Count the TP users
	filter = bson.M{
		"type_dni": domain.TypeDniTP,
	}
	total_TP_users, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return 0, 0, 0, err
	}
	return total_CC_users, total_TI_users, total_TP_users, nil

}
