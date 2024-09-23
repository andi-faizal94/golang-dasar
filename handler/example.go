package handler

import (
	"fmt"
	"log"
	"net/http"

	"golang-dasar/example"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)



var validate *validator.Validate

func InitValidator(v *validator.Validate) {
    validate = v
}


 func HandlerGet (ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"name":"ical",
		"age":12,
	})
  }
  func HandlerGetAll(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"name":"ical",
		"age":12,
	})
  }
  func HandlerGetById(ctx *gin.Context){
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK,gin.H{
		"id":id,
	})

  }
  func HandleQuery(ctx *gin.Context){
	biodata := ctx.Query("biodata")

	ctx.JSON(http.StatusOK,gin.H{
		"biodata":biodata,
	})
  }

  func PostHandler(ctx *gin.Context){
	var example example.Example


	// parsing to JSON

	err := ctx.ShouldBindJSON(&example)

	if err != nil {
		log.Fatal(err)
		
		return
	}
	
	
	// validation
	err = validate.Struct(example)



    if err != nil {
		// tipe data slices
		errorMessages := []string {}

		for _,e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on field %s,condition : %s",e.Field(), e.ActualTag())
			errorMessages = append(errorMessages,errorMessage)

        ctx.JSON(http.StatusBadRequest, gin.H{ "error":errorMessages})
		}
        return
    }

	ctx.JSON(http.StatusOK,gin.H{
		"name":example.Name,
		"age":example.Age,
	})
  }