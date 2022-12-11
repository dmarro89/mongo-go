package main

import (
	"context"
	"fmt"
	"mongo-go/api"
	"mongo-go/database"
	"mongo-go/service"

	"github.com/gin-gonic/gin"
)

const (
	ready  = "/ready"
	health = "/health"
	user   = "user"
	users  = "users"
)

var userId = fmt.Sprintf("%s/:userId", user)

func main() {
	mongoDB := database.New()
	err := mongoDB.Connect()
	if err != nil {
		panic(fmt.Sprintf("Impossible to start Microservice without DB Connection - %v", err))
	}

	r := gin.Default()

	api := api.New(service.New(mongoDB))
	r.POST(user, api.Create())
	r.PUT(user, api.Update())
	r.DELETE(userId, api.DeleteByID())
	r.GET(users, api.Get())
	r.GET(userId, api.GetByID())

	defer mongoDB.Disconnect(context.Background())
	err = r.Run()
	if err != nil {
		panic(fmt.Sprintf("Unexpected error during microservice run - %v", err))
	}
}
