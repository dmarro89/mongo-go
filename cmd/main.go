package main

import (
	"context"
	"fmt"
	"mongo-go/api"
	"mongo-go/database"
	"mongo-go/service"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	userEndpoint  = "user"
	usersEndpoint = "users"
)

var userId = fmt.Sprintf("%s/:userId", userEndpoint)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*3))
	defer cancel()
	mongoDB := database.New()
	err := mongoDB.Connect(ctx)
	if err != nil {
		panic(fmt.Sprintf("Impossible to start Microservice without DB Connection - %v", err))
	}

	r := gin.Default()

	api := api.New(service.New(mongoDB))
	r.POST(userEndpoint, api.Create())
	r.PUT(userEndpoint, api.Update())
	r.DELETE(userId, api.DeleteByID())
	r.GET(usersEndpoint, api.Get())
	r.GET(userId, api.GetByID())

	defer mongoDB.Disconnect(ctx)
	err = r.Run()
	if err != nil {
		panic(fmt.Sprintf("Unexpected error during microservice run - %v", err))
	}
}
