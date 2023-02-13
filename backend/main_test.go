package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/nguyendt456/software-engineer-project/database"
	"github.com/nguyendt456/software-engineer-project/models"
)

type UserTestSuite struct {
	User models.User
	suite.Suite
	R *gin.Engine
}

func SendRequest(engine *gin.Engine, method string, url string, payload interface{}, response interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(payload)

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()
	engine.ServeHTTP(writer, req)

	json.Unmarshal(writer.Body.Bytes(), response)

	return writer
}

func (suite *UserTestSuite) SetupSuite() {
	suite.R = SetupRouter()
	suite.User = models.User{
		UserName: "usertest",
		Name:     "TEST",
		Password: "123456789",
		UserType: "janitor",
	}

	var response models.Response
	SendRequest(suite.R, "POST", "/registry", suite.User, &response)

	assert.Equal(suite.T(), 201, response.Status)
	assert.Equal(suite.T(), "Success", response.Message)
}

func (suite *UserTestSuite) Test1_CreateDuplicateUser() {
	suite.User.Name = "neyugN"

	var response models.Response
	SendRequest(suite.R, "POST", "/registry", suite.User, &response)

	assert.Equal(suite.T(), 400, response.Status)
	assert.Equal(suite.T(), "Duplicate username", response.Message)
}

func (suite *UserTestSuite) Test2_LoginWrongPassword() {
	suite.User.Password = "123456781"

	var response models.Response
	SendRequest(suite.R, "POST", "/login", suite.User, &response)

	assert.Equal(suite.T(), 400, response.Status)
	assert.Equal(suite.T(), "Wrong username or password", response.Message)
}

func (suite *UserTestSuite) Test3_LoginRightPassword() {
	suite.User.Password = "123456789"

	var response models.User
	writer := SendRequest(suite.R, "POST", "/login", suite.User, &response)

	suite.User = response

	assert.Equal(suite.T(), 200, writer.Code)
}

func (suite *UserTestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result := database.UserCollection.FindOneAndDelete(ctx, bson.D{
		{
			Key:   "username",
			Value: suite.User.UserName,
		},
	})

	assert.Equal(suite.T(), nil, result.Err())
}

func TestMain(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
