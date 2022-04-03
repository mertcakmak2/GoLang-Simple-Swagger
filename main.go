package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "go-swagger/docs"
	"go-swagger/model"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath  /api/v1
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.New()

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", AuthMiddleware(), FindUsers)
			users.GET(":id", AuthMiddleware(), FindUserById)
			users.POST("", AuthMiddleware(), AddUser)
			users.DELETE(":id", AuthMiddleware(), DeleteUserById)
		}
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// FindUserById godoc
// @Summary Get all users.
// @Tags users
// @Accept */*
// @Produce json
// @Success      200  {array}   model.User
// @Security BearerAuth
// @Router   /users [get]
func FindUsers(c *gin.Context) {

	users := []model.User{
		model.User{ID: 1, Username: "user1", Password: "password"},
		model.User{ID: 2, Username: "user2", Password: "password"},
		model.User{ID: 3, Username: "user3", Password: "password"},
	}
	c.JSON(http.StatusOK, users)
}

// FindUserById godoc
// @Summary Find User by id.
// @Param id path int  true  "User ID"
// @Tags users
// @Accept */*
// @Produce json
// @Success 200 {object} model.User
// @Security BearerAuth
// @Router  /users/{id} [get]
func FindUserById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	user := model.User{ID: id, Username: "mertcakmak", Password: "password"}
	c.JSON(http.StatusOK, user)
}

// AddUser godoc
// @Summary      Add an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "Add user"
// @Success      201      {object}  model.User
// @Security BearerAuth
// @Router       /users [post]
func AddUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "Unauthorization")
		return
	}

	user = model.User{ID: user.ID, Username: "saved_mertcakmak", Password: "saved_password"}
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary      Delete an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204
// @Security BearerAuth
// @Router       /users/{id} [delete]
func DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusNoContent, "deleted user: "+id)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")
		fmt.Println(authorization)
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, "Unauthorization")
			c.Abort()
			return
		}
	}
}
