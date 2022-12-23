package service

import (
	"context"
	"mongo-go/database"
	"mongo-go/models"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	goassert "gotest.tools/assert"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

const (
	testName    = "Davide"
	testSurname = "Test"
	testEmail   = "davide@test.com"
)

var operationName = strings.Join([]string{testName, testSurname}, ".")

func TestNew(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	service := New(database.NewClient(mt.Client))
	assert.NotNil(t, service)
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	New(nil)
}

func TestPing(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Test ping", func(mt *mtest.T) {
		service := New(database.NewClient(mt.Client))
		assert.Equal(t, service.Ping(context.TODO()) != nil, true)
	})
}

func TestInsert(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Test insert success and fail", func(mt *mtest.T) {
		service := New(database.NewClient(mt.Client))
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := service.Insert(context.TODO(), models.User{Name: testName})
		assert.Equal(t, err == nil, true)
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{}))
		err = service.Insert(context.TODO(), models.User{Name: testName})
		assert.Equal(t, err != nil, true)
	})
}

func TestUpdate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Test update success and fail", func(mt *mtest.T) {
		service := New(database.NewClient(mt.Client))
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := service.Update(context.TODO(), models.User{Id: primitive.NewObjectID(), Name: testName})
		assert.Equal(t, err == nil, true)
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{}))
		err = service.Update(context.TODO(), models.User{Id: primitive.NewObjectID(), Name: testName})
		assert.Equal(t, err != nil, true)
	})
}

func TestGetByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Test find one success and fail", func(mt *mtest.T) {
		service := New(database.NewClient(mt.Client))
		ID := primitive.NewObjectID()
		expectedUser := models.User{
			Id:      ID,
			Name:    testName,
			Surname: testSurname,
			Email:   testEmail,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, operationName, mtest.FirstBatch, bson.D{
			{Key: "id", Value: expectedUser.Id},
			{Key: "name", Value: expectedUser.Name},
			{Key: "surname", Value: expectedUser.Surname},
			{Key: "email", Value: expectedUser.Email},
		}))

		user, err := service.GetByID(context.TODO(), ID.Hex())
		assert.Equal(t, err == nil, true)
		goassert.DeepEqual(t, &expectedUser, user)
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{}))
		user, err = service.GetByID(context.TODO(), ID.Hex())
		assert.Equal(t, err != nil, true)
		assert.Equal(t, user == nil, true)
	})
}

func TestGet(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Test Get success and fail", func(mt *mtest.T) {
		service := New(database.NewClient(mt.Client))
		ID := primitive.NewObjectID()
		expectedUser := models.User{
			Id:      ID,
			Name:    testName,
			Surname: testSurname,
			Email:   testEmail,
		}

		ID1 := primitive.NewObjectID()

		expectedUser1 := models.User{
			Id:      ID1,
			Name:    "Raffaele",
			Surname: testSurname,
			Email:   "raffaele@test.com",
		}

		first := mtest.CreateCursorResponse(1, operationName, mtest.FirstBatch, bson.D{
			{Key: "id", Value: expectedUser.Id},
			{Key: "name", Value: expectedUser.Name},
			{Key: "surname", Value: expectedUser.Surname},
			{Key: "email", Value: expectedUser.Email},
		})
		second := mtest.CreateCursorResponse(1, operationName, mtest.NextBatch, bson.D{
			{Key: "id", Value: expectedUser1.Id},
			{Key: "name", Value: expectedUser1.Name},
			{Key: "surname", Value: expectedUser1.Surname},
			{Key: "email", Value: expectedUser1.Email},
		})

		killCursors := mtest.CreateCursorResponse(0, operationName, mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		users, err := service.Get(context.TODO())
		assert.Equal(t, err == nil, true)
		assert.Equal(t, len(users), 2)
		goassert.DeepEqual(t, expectedUser, users[0])
		goassert.DeepEqual(t, expectedUser1, users[1])
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{}))
		users, err = service.Get(context.TODO())
		assert.Equal(t, err != nil, true)
		assert.Equal(t, users == nil, true)
	})
}

func TestDeleteByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Test Delete success and fail", func(mt *mtest.T) {
		service := New(database.NewClient(mt.Client))
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := service.DeleteByID(context.TODO(), primitive.NewObjectID().Hex())
		assert.Equal(t, err == nil, true)
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{}))
		err = service.DeleteByID(context.TODO(), primitive.NewObjectID().Hex())
		assert.Equal(t, err != nil, true)
	})
}
