package example

import "encoding/json"


type Example struct {
    Name string        `json:"name" validate:"required"`   
	Age  json.Number    `json:"age" binding:"required,min=1"`
}