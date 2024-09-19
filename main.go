package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Example struct {
    Name string `json:"name" validate:"required"`   
    Age  int    `json:"age" validate:"required,number"` 
}
var validate *validator.Validate

  
  func main() {
	router := gin.Default()
	validate = validator.New()

	router.GET("/example",handlerGetAll)
	router.GET("/",handlerGet)
	router.GET("/example/:id",handlerGetById)
	router.GET("/query",handleQuery)
	router.POST("/example",postHandler)
	router.Run() 
  }

  func handlerGet (ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"name":"ical",
		"age":12,
	})
  }
  func handlerGetAll(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"name":"ical",
		"age":12,
	})
  }
  func handlerGetById(ctx *gin.Context){
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK,gin.H{
		"id":id,
	})

  }
  func handleQuery(ctx *gin.Context){
	biodata := ctx.Query("biodata")

	ctx.JSON(http.StatusOK,gin.H{
		"biodata":biodata,
	})
  }

  func postHandler(ctx *gin.Context){
	var example Example

	err := ctx.ShouldBindJSON(&example)
	if err != nil {
		log.Fatal(err)
		return
	}

	// validation

	err = validate.Struct(example)

	// parsing to JSON

	fmt.Println(ctx.ShouldBindJSON(&example))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

	ctx.JSON(http.StatusOK,gin.H{
		"name":example.Name,
		"age":example.Age,
	})
  }