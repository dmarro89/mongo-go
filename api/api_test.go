package api

import (
	"bytes"
	"context"
	"encoding/json"
	"mongo-go/database"
	"mongo-go/models"
	"mongo-go/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	testName    = "Davide"
	testSurname = "Test"
	testEmail   = "davide@test.com"
)

var (
	testUser = models.User{Name: testName, Surname: testSurname, Email: testEmail}
)

func TestNew(t *testing.T) {
	mongo := database.New()
	mongo.Connect(context.Background())
	api := New(service.NewMock())
	assert.Equal(t, api != nil, true)
}

func TestCreate(t *testing.T) {
	api := New(service.NewMock())
	r := gin.Default()
	r.POST("/user", api.Create())
	jsonValue, _ := json.Marshal(testUser)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	jsonValue, _ = json.Marshal(map[string]string{"test": "data"})
	req, _ = http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdate(t *testing.T) {
	api := New(service.NewMock())
	r := gin.Default()
	r.PUT("/user", api.Update())
	user := models.User{
		Id:      primitive.NewObjectID(),
		Name:    "Davide",
		Surname: testSurname,
		Email:   "davide@test.com",
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("PUT", "/user", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	jsonValue, _ = json.Marshal(map[string]string{"test": "data"})
	req, _ = http.NewRequest("PUT", "/user", bytes.NewBuffer(jsonValue))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDelete(t *testing.T) {
	api := New(service.NewMock())
	r := gin.Default()
	r.DELETE("/user/:userId", api.DeleteByID())
	req, _ := http.NewRequest("DELETE", "/user/"+primitive.NewObjectID().Hex(), nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetByID(t *testing.T) {
	api := New(service.NewMock())
	r := gin.Default()
	r.GET("/user/:userId", api.GetByID())
	id := primitive.NewObjectID()
	req, _ := http.NewRequest("GET", "/user/"+id.Hex(), bytes.NewBuffer(nil))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGet(t *testing.T) {
	api := New(service.NewMock())
	r := gin.Default()
	r.GET("/users", api.Get())
	req, _ := http.NewRequest("GET", "/users", bytes.NewBuffer(nil))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
