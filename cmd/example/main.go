package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muchrief/gin_api"
)

type BaseResponse[T interface{}] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type UserData struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func RegisterApi(sdk *gin_api.GinApiSdk) {
	sdk.Register(&gin_api.ApiData{
		Method:       http.MethodGet,
		RelativePath: "/health",
		Response:     &BaseResponse[interface{}]{},
		Tags:         []string{"status"},
	}, func(ctx *gin.Context) {
		response := &BaseResponse[interface{}]{
			Message: "OK",
		}
		ctx.JSON(http.StatusOK, response)
	})

	sdk.Register(&gin_api.ApiData{
		Method:       http.MethodGet,
		RelativePath: "/user",
		Response:     &BaseResponse[UserData]{},
		Tags:         []string{"user"},
	}, func(ctx *gin.Context) {
		response := &BaseResponse[UserData]{
			Message: "OK",
			Data: UserData{
				ID:       1,
				Name:     "John Doe",
				Username: "john_doe",
				Email:    "johndoe@gmail.com",
				Phone:    "0812345XXXX",
			},
		}
		ctx.JSON(http.StatusOK, response)
	})
}

func main() {
	r := gin.Default()

	sdk := gin_api.NewGinApiSdk(r)
	doc := SetUpDocumentation(r)

	sdk.Use(func(a gin_api.Api) {
		doc.RegisterDoc(a)
	})

	RegisterApi(sdk)

	port := "localhost:3000"
	err := sdk.R.Run(port)
	if err != nil {
		panic(err)
	}
}
