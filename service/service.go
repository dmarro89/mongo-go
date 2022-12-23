package service

import (
	"context"
	"mongo-go/database"
	"mongo-go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	users   = "users"
	name    = "name"
	surname = "surname"
	email   = "email"
	id      = "id"
)

type Service interface {
	Ping(ctx context.Context) error
	Insert(ctx context.Context, user models.User) error
	Update(ctx context.Context, user models.User) error
	GetByID(ctx context.Context, userId string) (*models.User, error)
	Get(ctx context.Context) ([]models.User, error)
	DeleteByID(ctx context.Context, userId string) error
}

type UserService struct {
	mongoDB *database.MongoDB
}

func New(mongoDB *database.MongoDB) Service {
	if mongoDB == nil {
		panic(database.ErrClientNotInitialized)
	}
	return &UserService{mongoDB: mongoDB}
}

func (s *UserService) Ping(ctx context.Context) error {
	return s.mongoDB.GetClient().Ping(ctx, nil)
}

func (s *UserService) Insert(ctx context.Context, user models.User) error {
	collection := s.mongoDB.GetCollection(database.Database, users)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Update(ctx context.Context, user models.User) error {
	collection := s.mongoDB.GetCollection(database.Database, users)
	objId, err := primitive.ObjectIDFromHex(user.Id.Hex())
	if err != nil {
		return err
	}
	update := bson.M{name: user.Name, surname: user.Surname, email: user.Email}
	_, err = collection.UpdateOne(ctx, bson.M{id: objId}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetByID(ctx context.Context, userId string) (*models.User, error) {
	collection := s.mongoDB.GetCollection(database.Database, users)
	var user models.User
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, bson.M{id: objId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Get(ctx context.Context) ([]models.User, error) {
	collection := s.mongoDB.GetCollection(database.Database, users)
	var users []models.User

	res, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)

	if err = res.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) DeleteByID(ctx context.Context, userId string) error {
	collection := s.mongoDB.GetCollection(database.Database, users)
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{id: objId})
	if err != nil {
		return err
	}
	return nil
}
