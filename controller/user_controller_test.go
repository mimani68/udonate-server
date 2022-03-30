package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"udonate/entity"
	"udonate/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserController_Create(t *testing.T) {
	UserRepository.DeleteAll()
	createUserRequest := model.CreateUserRequest{
		Name:     "Test User",
		Price:    1000,
		Quantity: 1000,
	}
	requestBody, _ := json.Marshal(createUserRequest)

	request := httptest.NewRequest("POST", "/api/Users", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	createUserResponse := model.CreateUserResponse{}
	json.Unmarshal(jsonData, &createUserResponse)
	assert.NotNil(t, createUserResponse.Id)
	assert.Equal(t, createUserRequest.Name, createUserResponse.Name)
	assert.Equal(t, createUserRequest.Price, createUserResponse.Price)
	assert.Equal(t, createUserRequest.Quantity, createUserResponse.Quantity)
}

func TestUserController_List(t *testing.T) {
	UserRepository.DeleteAll()
	User := entity.User{
		Id:       uuid.New().String(),
		Name:     "Sample User",
		Price:    1000,
		Quantity: 1000,
	}
	UserRepository.Insert(User)

	request := httptest.NewRequest("GET", "/api/Users", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.([]interface{})
	containsUser := false

	for _, data := range list {
		jsonData, _ := json.Marshal(data)
		getUserResponse := model.GetUserResponse{}
		json.Unmarshal(jsonData, &getUserResponse)
		if getUserResponse.Id == User.Id {
			containsUser = true
		}
	}

	assert.True(t, containsUser)
}
