package main

import (
	"context"
	"softwareIIbackend/internal/core/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers(ctx context.Context, coll *mongo.Collection) error {
	users := []interface{}{
		domain.User{
			TypeDNI:   domain.TypeDniCC,
			DNI:       "12345",
			FirstName: "Admin",
			LastName:  "Adminstrator",
			Email:     "santiago.bedoya35419@ucaldas.edu.co",
			Password:  encryptPassword("admin12345"),
			Role:      domain.AdminRole,
		},
	}

	_, err := coll.InsertMany(ctx, users)
	return err
}

func encryptPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
