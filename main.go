package main

import (
	"golang-dasar/handler"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	// "github.com/go-playground/validator"
)

var validate *validator.Validate

  
  func main() {
	router := gin.Default()
	validate = validator.New()
	v1 := router.Group("/v1")
	handler.InitValidator(validate)



	v1.GET("/example",handler.HandlerGetAll)
	v1.GET("/",handler.HandlerGet)
	v1.GET("/example/:id",handler.HandlerGetById)
	v1.GET("/query",handler.HandleQuery)
	v1.POST("/example",handler.PostHandler)
	router.Run() 
  }

  