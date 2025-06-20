package go_doc_test

import (
	"testing"

	"github.com/muchrief/go_doc"
	"github.com/stretchr/testify/assert"
)

type ApiError struct {
	Message           string            `json:"message,omitempty"`
	ValidationMessage map[string]string `json:"validation_message"`
}

type ApiResponse[T any] struct {
	*ApiError
	Data T `json:"data"`
}

type User struct {
	Email    string `json:"email" fmt:"email" binding:"required,email"`
	Password string `json:"password" fmt:"password" binding:"required,password"`
	Username string `json:"username" binding:"required,gte=6,lte=32"`
}

type ListUser []*User

type UserMapObject struct {
	UserMapper map[string]*User `json:"user_mapper"`
}

type PageInfo struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
type Response struct {
	Message  string    `json:"message"`
	PageInfo *PageInfo `json:"page_info"`
}

func TestBuildSchema(t *testing.T) {

	t.Run("test not pointer", func(t *testing.T) {
		apiError := ApiError{}
		result := go_doc.BuildSchema(apiError)
		assert.NotEmpty(t, result)
		assert.NotEmpty(t, result.Properties)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test not pointer with pointer", func(t *testing.T) {
		apiError := &ApiError{}
		result := go_doc.BuildSchema(apiError)
		assert.NotEmpty(t, result)
		assert.NotEmpty(t, result.Properties)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test embed struct and generic", func(t *testing.T) {
		response := &ApiResponse[map[string]string]{}
		result := go_doc.BuildSchema(response)
		assert.NotEmpty(t, result)
		assert.NotEmpty(t, result.Properties)
		assert.NotEmpty(t, result.Properties["validation_message"])
		assert.NotEmpty(t, result.Properties["data"].AdditionalProperties)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test pointer field", func(t *testing.T) {
		data := &Response{}
		result := go_doc.BuildSchema(data)
		assert.NotEmpty(t, result)
		assert.NotEmpty(t, result.Properties)
		assert.NotEmpty(t, result.Properties["page_info"])

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test fmt tag", func(t *testing.T) {
		usr := &User{}
		result := go_doc.BuildSchema(usr)
		assert.NotEmpty(t, result)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test map to struct", func(t *testing.T) {
		usr := &UserMapObject{}
		result := go_doc.BuildSchema(usr)
		assert.NotEmpty(t, result)
		assert.NotEmpty(t, result.Properties["user_mapper"])
		assert.NotEmpty(t, result.Properties["user_mapper"].AdditionalProperties)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test array", func(t *testing.T) {
		data := ListUser{}
		result := go_doc.BuildSchema(data)
		assert.NotEmpty(t, result)
		assert.NotEmpty(t, result.Items)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})
}
