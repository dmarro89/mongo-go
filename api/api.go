package api

import (
	"log"
	"mongo-go/models"
	"mongo-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API interface {
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	DeleteByID() gin.HandlerFunc
	GetByID() gin.HandlerFunc
	Get() gin.HandlerFunc
}

const (
	userId   = "userId"
	errorStr = "error"
)

type UserAPI struct {
	service *service.UserService
}

func New(service *service.UserService) API {
	return &UserAPI{service: service}
}

func (api *UserAPI) Create() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user models.User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{errorStr: err.Error()})
			return
		}
		err = api.service.Insert(c.Request.Context(), user)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{errorStr: err.Error()})
			return
		}
		c.String(http.StatusOK, "User created")
	}

	return gin.HandlerFunc(fn)
}

func (api *UserAPI) Update() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user models.User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{errorStr: err.Error()})
			return
		}
		err = api.service.Update(c.Request.Context(), user)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{errorStr: err.Error()})
			return
		}
		c.String(http.StatusOK, "User updated")
	}

	return gin.HandlerFunc(fn)
}

func (api *UserAPI) DeleteByID() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userId := c.Param(userId)
		err := api.service.DeleteByID(c.Request.Context(), userId)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{errorStr: err.Error()})
			return
		}
		c.String(http.StatusOK, "User deleted")
	}

	return gin.HandlerFunc(fn)
}

func (api *UserAPI) GetByID() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userId := c.Param(userId)
		user, err := api.service.GetByID(c.Request.Context(), userId)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{errorStr: err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	}

	return gin.HandlerFunc(fn)
}

func (api *UserAPI) Get() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		users, err := api.service.Get(c.Request.Context())
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{errorStr: err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"users": users})
	}

	return gin.HandlerFunc(fn)
}
